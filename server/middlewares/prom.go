package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/prom"
	"github.com/kevin-luvian/goauth/server/pkg/util"
)

// add the middleware function
func HttpMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		logging.Infoln("log metrics", c.Request.URL.String())

		c.Next()

		prom.CollectHttp(prom.HTTPMetrics{
			ClientIP:     util.GetClientIPAddr(c.Request),
			Method:       c.Request.Method,
			Route:        c.Request.URL.String(),
			StatusCode:   c.Writer.Status(),
			Duration:     float64(time.Since(start).Milliseconds()) / 1000,
			RequestSize:  float64(util.ComputeApproximateRequestSize(c.Request)),
			ResponseSize: float64(c.Writer.Size()),
		})
	}
}
