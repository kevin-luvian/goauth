package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpLabels     = []string{"method", "code", "route"}
	functionLabels = []string{"fname", "file", "module"}
)

var requestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Tracks the number of HTTP requests.",
	},
	[]string{"method", "code", "route", "ip"})

var requestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Tracks the latencies for HTTP requests.",
		Buckets: []float64{0.1, 0.15, 0.25, 0.5, 0.7},
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

var functionsDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "functions_duration_microseconds",
		Help:    "Tracks the duration for functions calls in microseconds.",
		Buckets: []float64{1000, 2500, 5000, 7000, 10000, 20000},
	},
	functionLabels,
)
