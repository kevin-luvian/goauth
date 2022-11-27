package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpLabels = []string{"method", "code", "route"}
)

var requestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Tracks the number of HTTP requests.",
	},
	httpLabels)

var requestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Tracks the latencies for HTTP requests.",
		Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
	},
	httpLabels,
)

var requestSize = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_size_bytes",
		Help: "Tracks the size of HTTP requests.",
	},
	httpLabels,
)

var responseSize = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_response_size_bytes",
		Help: "Tracks the size of HTTP responses.",
	},
	httpLabels,
)
