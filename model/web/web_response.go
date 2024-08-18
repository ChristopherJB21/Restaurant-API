package web

import "net/http"

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseStatus struct {
	http.ResponseWriter
	Status int
}

func (statusResponse *ResponseStatus) WriteHeader(code int) {
	statusResponse.ResponseWriter.WriteHeader(code)
	statusResponse.Status = code
}
