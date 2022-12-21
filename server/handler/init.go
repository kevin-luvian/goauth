package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/app"
	"github.com/kevin-luvian/goauth/server/usecases"
)

type Handler struct {
	authUC usecases.IAuthUseCase
}

type Dependencies struct {
	AuthUC usecases.IAuthUseCase
}

func New(dep Dependencies) *Handler {
	return &Handler{
		authUC: dep.AuthUC,
	}
}

func (h *Handler) HandlerPing(r gin.IRoutes) gin.IRoutes {
	return r.GET("/", func(c *gin.Context) {
		app.Success(c, map[string]interface{}{
			"ok": true,
		})
	})
}

func setRefreshTokenCookie(c *gin.Context, refreshToken string) {
	c.SetCookie("refresh-token", refreshToken, 0, "/", "", true, true)
}

func getRefreshTokenCookie(c *gin.Context) string {
	tok, err := c.Cookie("refresh-token")
	if err != nil {
		return ""
	}

	return tok
}
