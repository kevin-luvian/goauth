package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/handler"
	"github.com/kevin-luvian/goauth/server/pkg/db"
	"github.com/kevin-luvian/goauth/server/pkg/gconsul"
	"github.com/kevin-luvian/goauth/server/pkg/goagain"
	"github.com/kevin-luvian/goauth/server/pkg/gredis"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/prom"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
	"github.com/kevin-luvian/goauth/server/pkg/util"
	authRepo "github.com/kevin-luvian/goauth/server/repositories/auth"
	userRepo "github.com/kevin-luvian/goauth/server/repositories/user"
	"github.com/kevin-luvian/goauth/server/routers"
	authUC "github.com/kevin-luvian/goauth/server/usecases/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	util.Setup()
	setting.Setup()
	logging.Setup()

	err := gconsul.Setup()
	if err != nil {
		logging.Errorln("Error consul setup", err)
	}

	err = gconsul.FetchKV()
	if err != nil {
		logging.Errorln("Error consul fetch", err)
	}

	prom.Setup()
	gredis.Setup()
}

func main() {
	gin.SetMode(setting.Server.RunMode)

	// setup dependencies config
	ga := &oauth2.Config{
		ClientID:     setting.GoogleOAuth.ClientID,
		ClientSecret: setting.GoogleOAuth.SecretID,
		RedirectURL:  setting.GoogleOAuth.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	// setup database
	db, err := db.New(db.Config{
		SourceURL:             setting.Database.DockerURL,
		Retries:               setting.Database.Retries,
		MaxOpenConnections:    setting.Database.MaxActive,
		MaxIdleConnections:    setting.Database.MaxIdle,
		ConnectionMaxLifetime: setting.Database.MaxLifetime,
	})
	if err != nil {
		logging.Fatalln("error initializing database", err)
	}

	// setup repositories
	authRepo := authRepo.New(ga)
	userRepo := userRepo.New(db)

	// setup usecases
	authUC := authUC.New(authUC.Dependencies{
		AuthRepo: authRepo,
		UserRepo: userRepo,
	})

	// setup handler
	h := handler.New(handler.Dependencies{
		AuthUC: authUC,
	})

	routersInit := routers.InitRouter(h)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Server.HttpPort),
		Handler:        routersInit,
		ReadTimeout:    setting.Server.ReadTimeout,
		WriteTimeout:   setting.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// Use graceful restart and get listener from parent
	l, err := goagain.Listener()
	if nil != err {
		l, err = net.Listen("tcp", server.Addr)
		if nil != err {
			logging.Fatalln(err)
		}

		logging.Infoln("listening on", server.Addr)
		go server.Serve(l)

	} else {
		// Resume accepting connections in a new goroutine.
		logging.Infoln("resuming listening on", server.Addr)
		go server.Serve(l)

		// Kill the parent, now that the child has started successfully.
		if err := goagain.Kill(); nil != err {
			logging.Fatalln(err)
		}
	}

	runTicker(l)

	// Block the main goroutine awaiting signals.
	if _, err := goagain.Wait(l); nil != err {
		logging.Fatalln(err)
	}

	// Do whatever's necessary to ensure a graceful exit like waiting for
	// goroutines to terminate or a channel to become closed.
	if err := l.Close(); nil != err {
		logging.Fatalln(err)
	}

	logging.Infoln("closing http server on", server.Addr)
	time.Sleep(1 * time.Second)
}

func runTicker(l net.Listener) {
	// periodic updates
	goagain.Ticker(setting.App.TickerTTL, func() {
		start := time.Now().UnixMilli()

		gconsul.NotifyHealth(func() error {
			return gredis.Ping()
		})

		if gconsul.HasKVChanged() || setting.HasSettingChanged() {
			logging.Infoln("configuration changes detected, restarting...")

			err := goagain.ForkExec(l)
			if nil != err {
				logging.Errorln(err)
			}
		}

		elapsed := time.Now().UnixMilli() - start
		logging.Infof("ticking... %dms", elapsed)
	})
}
