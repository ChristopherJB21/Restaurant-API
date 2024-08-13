package middleware

import (
	"crypto/rsa"
	"net/http"
	"restaurant/helper"
	"restaurant/model/web"
)

type Middleware struct {
	Handler      http.Handler
	RSAPublicKey *rsa.PublicKey
}

func NewMiddleware(handler http.Handler, rSAPublicKey *rsa.PublicKey) *Middleware {
	return &Middleware{
		Handler:      handler,
		RSAPublicKey: rSAPublicKey,
	}
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if middleware.Authorization(writer, request){
		return
	}

	middleware.Handler.ServeHTTP(writer, request)
}

func (middleware *Middleware) Authorization(writer http.ResponseWriter, request *http.Request) bool {
	err := helper.VerifyToken(request, middleware.RSAPublicKey)

	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   err.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	}

	return false
}