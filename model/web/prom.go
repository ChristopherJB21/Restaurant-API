package web

import "github.com/prometheus/client_golang/prometheus"

type MetricPrometheus struct {
	RequestCounter    *prometheus.CounterVec
	DurationHistogram *prometheus.HistogramVec
}
