package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/app"
)

func GetAuth(c *gin.Context) {
	app.Success(c, map[string]interface{}{
		"ok": true,
	})
}
