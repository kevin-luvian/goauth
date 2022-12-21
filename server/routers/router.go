package routers

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/handler"
	"github.com/kevin-luvian/goauth/server/middlewares"
	"github.com/kevin-luvian/goauth/server/pkg/prom"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
)

// InitRouter initialize routing information
func InitRouter(h *handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// r.Use(gin.Logger())
	// r.Use(func(c *gin.Context) {
	// 	or := c.GetHeader("origin")
	// 	logging.Infoln("Header Origin", or)
	// })

	corsRules := CreateCORSRule(strings.Split(setting.App.CORS, ";"))
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: CheckOrigin(corsRules),
		AllowHeaders:    []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin"},
		AllowMethods:    []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
	}))

	root := r.Group("/api")
	{
		root.GET("/metrics", prom.Handler())
		h.HandlerPing(root)
	}

	apiv1 := r.Group("/api/v1", middlewares.HttpMetricsMiddleware())
	{
		apiAuth := apiv1.Group("/auth")
		{
			h.HandlerAuthPing(apiAuth)
			h.HandlerGoogleLogin(apiAuth)
			h.HandlerGoogleSignup(apiAuth)
			h.HandlerAuthenticateGoogleRedirectOrigin(apiAuth)
		}
	}

	return r
}
