package controller

import (
	"crypto/rsa"
	"net/http"
	"restaurant/helper"
	model "restaurant/model/cuisine"
	"restaurant/model/web"
	"restaurant/service"
	"strconv"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

type ICuisineController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type CuisineController struct {
	CuisineService service.ICuisineService
	RSAPublicKey *rsa.PublicKey
}

func NewCuisineController(cuisineService service.ICuisineService, rSAPublicKey *rsa.PublicKey) ICuisineController {
	return &CuisineController{
		CuisineService: cuisineService,
		RSAPublicKey: rSAPublicKey,
	}
}

func (controller *CuisineController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cuisineCreateRequest := model.CuisineCreateRequest{}
	var result interface{} = &cuisineCreateRequest
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	helper.PanicIfError(err)

	username := helper.GetUsername(request, controller.RSAPublicKey)
	cuisineCreateRequest.CreatedBy = username
	cuisineCreateRequest.UpdatedBy = username

	cuisineResponse := controller.CuisineService.Create(request.Context(), cuisineCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   cuisineResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CuisineController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	IDCuisine := params.ByName("IDCuisine")
	id, err := strconv.ParseUint(IDCuisine, 10, 64)
	helper.PanicIfError(err)

	cuisineDeleteRequest := model.CuisineDeleteRequest{}
	cuisineDeleteRequest.IDCuisine = uint(id)
	cuisineDeleteRequest.DeletedBy = helper.GetUsername(request, controller.RSAPublicKey)

	controller.CuisineService.Delete(request.Context(), cuisineDeleteRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CuisineController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var limit, offset int

	limitQuery, err := helper.ReadFromQueryParams("limit", request)
	if err != nil {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitQuery)
		helper.PanicIfError(err)
	}

	offsetQuery, err := helper.ReadFromQueryParams("offset", request)
	if err != nil {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetQuery)
		helper.PanicIfError(err)
	}

	cuisineResponses := controller.CuisineService.FindAll(request.Context(), limit, offset)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   cuisineResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CuisineController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	IDCuisine := params.ByName("IDCuisine")
	id, err := strconv.ParseUint(IDCuisine, 10, 64)
	helper.PanicIfError(err)

	cuisineResponse := controller.CuisineService.FindById(request.Context(), uint(id))

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   cuisineResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CuisineController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cuisineUpdateRequest := model.CuisineUpdateRequest{}
	helper.ReadFromRequestBody(request, &cuisineUpdateRequest)

	IDCuisine := params.ByName("IDCuisine")
	id, err := strconv.ParseUint(IDCuisine, 10, 64)
	helper.PanicIfError(err)

	cuisineUpdateRequest.IDCuisine = uint(id)

	username := helper.GetUsername(request, controller.RSAPublicKey)
	cuisineUpdateRequest.UpdatedBy = username

	cuisineResponse := controller.CuisineService.Update(request.Context(), cuisineUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   cuisineResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
