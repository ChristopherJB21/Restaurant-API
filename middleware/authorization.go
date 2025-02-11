package middleware

import (
	"net/http"
	"restaurant/helper"
	"restaurant/model/web"

	"github.com/spf13/viper"
)

func (middleware *Middleware) Authorization(writer http.ResponseWriter, request *http.Request) bool {
	DecodedAPIKey := GetAppKey(request)

	if DecodedAPIKey != viper.GetString("apiKey") {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
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
