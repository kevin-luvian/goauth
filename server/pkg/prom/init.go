package prom

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
)

var registry *prometheus.Registry

func Setup() {
	var identifier = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "goauth",
		Name:      "process_identifier",
	}, []string{"hostname"})
	identifier.WithLabelValues("goauth_instance").Inc()

	registry = prometheus.NewRegistry()
	registry.MustRegister(
		identifier,
		requestsTotal,
		requestDuration,
		requestSize,
		responseSize,
		version.NewCollector("version"),
	)
}

func Handler() gin.HandlerFunc {
	h := promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: false,
		})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
