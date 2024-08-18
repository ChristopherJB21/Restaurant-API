package middleware

import (
	"net/http"
	"restaurant/helper"
	"restaurant/model/web"
)

func (middleware *Middleware) Authorization(writer http.ResponseWriter, request *http.Request) bool {
	err := helper.VerifyToken(request, middleware.RSAPublicKey)

	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   err.Error(),
		}

		helper.WriteToResponseBody(writer, http.StatusUnauthorized, webResponse)

		return true
	}

	return false
}
