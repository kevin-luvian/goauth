package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/handler"
	"github.com/kevin-luvian/goauth/server/pkg/gredis"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
	"github.com/kevin-luvian/goauth/server/pkg/util"
	authRepo "github.com/kevin-luvian/goauth/server/repositories/auth"
	"github.com/kevin-luvian/goauth/server/routers"
	authUC "github.com/kevin-luvian/goauth/server/usecases/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	setting.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	ga := &oauth2.Config{
		ClientID:     setting.GoogleOAuthSetting.ClientID,
		ClientSecret: setting.GoogleOAuthSetting.SecretID,
		RedirectURL:  setting.GoogleOAuthSetting.RedirectURL,
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

	// swagger.GenerateSwaggerDocsAndEndpoints(routersInit.,"asd")
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	logging.Infof("start http server listening on %s", endPoint)

	server.ListenAndServe()
}
