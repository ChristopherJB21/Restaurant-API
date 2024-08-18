package web

import "net/http"

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type RequestStatus struct {
	http.ResponseWriter
	Status int
}

func (statusRequest *RequestStatus) WriteHeader(code int) {
	statusRequest.ResponseWriter.WriteHeader(code)
	statusRequest.Status = code
}
