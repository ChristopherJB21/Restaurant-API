package middleware

import (
	"net/http"
	"restaurant/helper"
	"restaurant/model/web"

	"github.com/spf13/viper"
)

type Middleware struct {
	Handler      http.Handler
}

func NewMiddleware(handler http.Handler) *Middleware {
	return &Middleware{
		Handler:      handler,
	}
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Call Middleware Authorization
	if middleware.Authorization(writer, request){
		return
	}

	middleware.Handler.ServeHTTP(writer, request)
}

func (middleware *Middleware) Authorization(writer http.ResponseWriter, request *http.Request) bool {
	// Get APIKey
	APIKey := request.Header.Get("X-API-Key")

	if APIKey != viper.GetString("apiKey") {
		// IF UNAUTHORIZED
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, http.StatusUnauthorized, webResponse)

		return true
	}

	return false
}