package web

import "github.com/prometheus/client_golang/prometheus"

type MetricPrometheus struct {
	RequestGauge      *prometheus.GaugeVec
	DurationHistogram *prometheus.HistogramVec
}
