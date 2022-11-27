package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/app"
	"github.com/kevin-luvian/goauth/server/pkg/e"
)

func Ping(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"ok": true,
	})
}

func PingBad(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusInternalServerError, e.ERROR, gin.H{
		"ok": false,
	})
}
