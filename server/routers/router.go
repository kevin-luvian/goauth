package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kevin-luvian/goauth/server/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tagss", v1.GetTags)
	}

	return r
}
