package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/handler"
	"github.com/kevin-luvian/goauth/server/pkg/gconsul"
	"github.com/kevin-luvian/goauth/server/pkg/goagain"
	"github.com/kevin-luvian/goauth/server/pkg/gredis"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/prom"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
	authRepo "github.com/kevin-luvian/goauth/server/repositories/auth"
	"github.com/kevin-luvian/goauth/server/routers"
	authUC "github.com/kevin-luvian/goauth/server/usecases/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	setting.Setup()
	logging.Setup()

	err := gconsul.Setup(func() error {
		return nil
	})
	if err != nil {
		logging.Fatalln("Error consul setup", err)
	}

	err = gconsul.FetchKV()
	if err != nil {
		logging.Fatalln("Error consul fetch", err)
	}

	prom.Setup()
	gredis.Setup()
	if err := gredis.Ping(); err != nil {
		logging.Fatalln("Error redis setup", err)
	}
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

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

	authRepo := authRepo.New(ga)

	authUC := authUC.New(authUC.Dependencies{
		AuthRepo: authRepo,
	})

	h := handler.New(handler.Dependencies{
		AuthUC: authUC,
	})

	routersInit := routers.InitRouter(h)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        routersInit,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// Use graceful restart and get listener from parent
	l, err := goagain.Listener()
	if nil != err {
		l, err = net.Listen("tcp", server.Addr)
		if nil != err {
			logging.Fatalln(err)
		}

		logging.Infoln("listening on", l.Addr())

		// Accept connections in a new goroutine.
		go server.Serve(l)

	} else {
		// Resume accepting connections in a new goroutine.
		logging.Infoln("resuming listening on", l.Addr())
		go server.Serve(l)

		// Kill the parent, now that the child has started successfully.
		if err := goagain.Kill(); nil != err {
			logging.Fatalln(err)
		}
	}

	// watch consul changes
	gconsul.WatchKV(func() {
		logging.Infoln("changes on consul detected, restarting...")

		err := goagain.ForkExec(l)
		if nil != err {
			logging.Errorln(err)
		}
	})

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
