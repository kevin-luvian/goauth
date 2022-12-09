package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/app"
)

func (h *Handler) HandlerAuthPing(c *gin.Context) {
	app.Success(c, map[string]interface{}{
		"ok": true,
	})
}
