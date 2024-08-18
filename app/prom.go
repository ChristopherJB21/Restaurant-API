package app

import (
	"net/http"
	"restaurant/helper"
	"restaurant/model/web"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

func NewMetricPrometheus() *web.MetricPrometheus {
	metrics := web.MetricPrometheus{
		RequestGauge:      NewRequestGauge(),
		DurationHistogram: NewDurationHistogram(),
	}

	return &metrics
}

func StartPrometheus() {
	router := httprouter.New()

	router.Handler("GET", "/metrics", promhttp.Handler())

	server := http.Server{
		Addr:    viper.GetString("server.prom"),
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

func NewRequestGauge() *prometheus.GaugeVec {
	requestCounter := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "client_request_count",
			Help: "Total number of requests from client",
		},
		[]string{"method", "route"},
	)

	return requestCounter
}

func NewDurationHistogram() *prometheus.HistogramVec {
	durationHistogram := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "client_request_duration_secs",
			Help: "Duration of requests from Client",
		},
		[]string{"method", "route", "status"},
	)

	return durationHistogram
}
