package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/prom"
)

// add the middleware function
func HttpMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		prom.CollectHttp(prom.HTTPMetrics{
			ClientIP:     getClientIPAddr(c.Request),
			Method:       c.Request.Method,
			Route:        c.Request.URL.String(),
			StatusCode:   c.Writer.Status(),
			Duration:     float64(time.Since(start).Milliseconds()) / 1000,
			RequestSize:  float64(computeApproximateRequestSize(c.Request)),
			ResponseSize: float64(c.Writer.Size()),
		})
	}
}

func getClientIPAddr(req *http.Request) string {
	ipSlice := []string{
		req.Header.Get("X-FORWARDED-FOR"),
		req.Header.Get("X-Forwarded-For"),
		req.Header.Get("x-forwarded-for"),
	}

	for _, ip := range ipSlice {
		if ip != "" {
			return ip
		}
	}

	return strings.Split(req.RemoteAddr, ":")[0]
}

func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s += len(r.URL.String())
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}
