package exception

import (
	"restaurant/helper"
	"restaurant/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	// Authentication Errror
	if authenticationError(writer, request, err) {
		return
	}

	// Not Found Error
	if notFoundError(writer, request, err) {
		return
	}

	// Validation Error
	if validationErrors(writer, request, err) {
		return
	}

	// Bad Request Error
	if badRequestError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func authenticationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(AuthenticationError)
	if ok {
		logger := helper.NewLogger()

		logger.WithFields(logrus.Fields{
			"status":        http.StatusUnauthorized,
			"error message": exception.Error,
		}).Warn("UNAUTHORIZED")

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, http.StatusUnauthorized, webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		logger := helper.NewLogger()

		logger.WithFields(logrus.Fields{
			"status":        http.StatusBadRequest,
			"error message": exception.Error,
		}).Warn("BAD REQUEST")

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, http.StatusBadRequest, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		logger := helper.NewLogger()

		logger.WithFields(logrus.Fields{
			"status":        http.StatusNotFound,
			"error message": exception.Error,
		}).Warn("NOT FOUND")

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, http.StatusNotFound, webResponse)
		return true
	} else {
		return false
	}
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		logger := helper.NewLogger()

		logger.WithFields(logrus.Fields{
			"status":        http.StatusBadRequest,
			"error message": exception.Error,
		}).Warn("BAD REQUEST")

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, http.StatusBadRequest, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	logger := helper.NewLogger()

	logger.WithFields(logrus.Fields{
		"status":        http.StatusInternalServerError,
		"error message": err,
	}).Error("INTERNAL SERVER ERROR")

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, http.StatusInternalServerError, webResponse)
}
