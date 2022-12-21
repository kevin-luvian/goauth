package prom

import (
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type HTTPMetrics struct {
	ClientIP     string
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

	requestsTotal.With(prometheus.Labels{
		"method": metrics.Method,
		"code":   strconv.Itoa(metrics.StatusCode),
		"route":  metrics.Route,
		"ip":     metrics.ClientIP,
	}).Inc()

	requestDuration.With(label).Observe(metrics.Duration)
	requestSize.With(label).Observe(metrics.RequestSize)
	responseSize.With(label).Observe(metrics.ResponseSize)
}

func CollectRepoDuration(funcName string) func() {
	return CollectFunctionDuration(funcName, "repository", 1)
}

func CollectFunctionDuration(funcName, module string, skip int) func() {
	label := prometheus.Labels{
		"fname":  funcName,
		"file":   getCallerFile(2 + skip),
		"module": module,
	}

	start := time.Now()
	return func() {
		elapsed := time.Since(start).Microseconds()
		functionsDuration.With(label).Observe(float64(elapsed))
	}
}

func getCallerFile(skip int) string {
	_, file, _, ok := runtime.Caller(skip)
	if ok {
		baseIndex := strings.LastIndex(file, "/server") + len("/server")
		if len(file) < baseIndex {
			return ""
		}

		file = file[baseIndex:]

		return file
	}

	return ""
}
