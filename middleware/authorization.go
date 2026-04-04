package middleware

import (
	"errors"
	"net/http"
	"restaurant/helper"
	"restaurant/model/web"

	"github.com/spf13/viper"
)

func (middleware *Middleware) Authorization(writer http.ResponseWriter, request *http.Request) bool {
	err := error(nil)

	if request.URL.Path == "/api/user/login" || (request.URL.Path == "/api/user" && request.Method == "POST") {
		apiKey := GetAppKey(request)

		if apiKey == viper.GetString("apiKey") {
			return false
		}

		err = errors.New("invalid API key")
	} else {
		err = helper.VerifyToken(request, middleware.RSAPublicKey)
	}

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

func GetAppKey(request *http.Request) string {
	var APIKey = request.Header.Get("X-API-Key")

	return APIKey
}
