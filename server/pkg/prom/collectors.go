package prom

import (
	"strconv"

	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/prometheus/client_golang/prometheus"
)

type HTTPMetrics struct {
	Method       string
	StatusCode   int
	Route        string
	Duration     float64
	RequestSize  float64
	ResponseSize float64
}

func CollectHttp(metrics HTTPMetrics) {
	label := prometheus.Labels{
		"method": metrics.Method,
		"code":   strconv.Itoa(metrics.StatusCode),
		"route":  metrics.Route,
	}

	logging.Info(label)

	requestsTotal.With(label).Inc()
	requestDuration.With(label).Observe(metrics.Duration)
	requestSize.With(label).Observe(metrics.RequestSize)
	responseSize.With(label).Observe(metrics.ResponseSize)
}
