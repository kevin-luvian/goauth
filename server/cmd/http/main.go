package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/gredis"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
	"github.com/kevin-luvian/goauth/server/pkg/util"
	"github.com/kevin-luvian/goauth/server/routers"
)

func init() {
	setting.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
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
