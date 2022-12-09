package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/handler"
	"github.com/kevin-luvian/goauth/server/middlewares"
	"github.com/kevin-luvian/goauth/server/pkg/prom"
	"github.com/kevin-luvian/goauth/server/routers/api"
	v1 "github.com/kevin-luvian/goauth/server/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter(h *handler.Handler) *gin.Engine {
	r := gin.New()
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())

	prom.Setup()

	root := r.Group("/api")
	{
		root.GET("/metrics", prom.Handler())
		root.GET("/ping", api.Ping)
		root.GET("/ping/bad", api.PingBad)
	}

	apiv1 := r.Group("/api/v1", middlewares.HttpMetricsMiddleware())
	{
		apiAuth := apiv1.Group("/auth")
		{
			apiAuth.GET("/ping", h.HandlerAuthPing)
		}

		apiv1.GET("/tags", v1.GetTags)
	}

	return r
}
