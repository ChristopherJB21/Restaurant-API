package middleware

import (
	"net/http"
	"restaurant/model/web"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

func (middleware *Middleware) PromMonitorBefore(writer web.ResponseStatus, request *http.Request) {
	var requestMethod string

	if len(request.Method) > 0 {
		requestMethod = request.Method
	} else {
		requestMethod = "GET"
	}

	labels := prometheus.Labels{
		"method": requestMethod,
		"route":  request.URL.Path,
	}

	middleware.MetricPrometheus.RequestGauge.With(labels).Inc()
}

func (middleware *Middleware) PromMonitorAfter(writer web.ResponseStatus, request *http.Request, duration float64) {
	var requestMethod string

	if len(request.Method) > 0 {
		requestMethod = request.Method
	} else {
		requestMethod = "GET"
	}

	labels := prometheus.Labels{
		"method": requestMethod,
		"route":  request.URL.Path,
		"status": strconv.Itoa(writer.Status),
	}

	middleware.MetricPrometheus.DurationHistogram.With(labels).Observe(duration)

	labels = prometheus.Labels{
		"method": requestMethod,
		"route":  request.URL.Path,
	}

	middleware.MetricPrometheus.RequestGauge.With(labels).Dec()
}
