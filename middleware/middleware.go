package middleware

import (
	"net/http"
	"restaurant/helper"
	"restaurant/model/web"

	"github.com/go-playground/validator/v10"
)

type Middleware struct {
	Handler  http.Handler
	Validate *validator.Validate
}

func NewMiddleware(handler http.Handler, validate *validator.Validate) *Middleware {
	return &Middleware{
		Handler:  handler,
		Validate: validate,
	}
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := helper.VerifyToken(request)

	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data: err.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return
	}

	middleware.Handler.ServeHTTP(writer, request)
}