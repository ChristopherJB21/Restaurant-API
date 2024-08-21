package middleware

import (
	"crypto/rsa"
	"net/http"
	"restaurant/model/web"
	"time"
)

type Middleware struct {
	Handler          http.Handler
	RSAPublicKey     *rsa.PublicKey
	MetricPrometheus *web.MetricPrometheus
}

func NewMiddleware(handler http.Handler, rSAPublicKey *rsa.PublicKey, metricPrometheus *web.MetricPrometheus) *Middleware {
	return &Middleware{
		Handler:          handler,
		RSAPublicKey:     rSAPublicKey,
		MetricPrometheus: metricPrometheus,
	}
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()

	if middleware.Authorization(writer, request) {
		return
	}

	ResponseStatus := web.ResponseStatus{ResponseWriter: writer, Status: 200}

	middleware.PromMonitorBefore(ResponseStatus, request)

	middleware.Handler.ServeHTTP(&ResponseStatus, request)

	duration := time.Since(start).Seconds()

	middleware.PromMonitorAfter(ResponseStatus, request, float64(duration))
}
