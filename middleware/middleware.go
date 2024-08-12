package middleware

import (
	"crypto/rsa"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Middleware struct {
	Handler  http.Handler
	Validate *validator.Validate
	RSAPublicKey *rsa.PublicKey
}

func NewMiddleware(handler http.Handler, validate *validator.Validate, rSAPublicKey *rsa.PublicKey) *Middleware {
	return &Middleware{
		Handler:  handler,
		Validate: validate,
		RSAPublicKey: rSAPublicKey,
	}
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// err := helper.VerifyToken(request, middleware.RSAPublicKey)

	// if err != nil {
	// 	writer.Header().Add("Content-Type", "application/json")
	// 	writer.WriteHeader(http.StatusUnauthorized)

	// 	webResponse := web.WebResponse{
	// 		Code:   http.StatusUnauthorized,
	// 		Status: "UNAUTHORIZED",
	// 		Data: err.Error(),
	// 	}

	// 	helper.WriteToResponseBody(writer, webResponse)

	// 	return
	// }

	middleware.Handler.ServeHTTP(writer, request)
}