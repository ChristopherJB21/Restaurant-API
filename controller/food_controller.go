package controller

import (
	"crypto/rsa"
	"net/http"
	"restaurant/helper"
	model "restaurant/model/food"
	"restaurant/model/web"
	"restaurant/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type IFoodController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type FoodController struct {
	FoodService  service.IFoodService
	RSAPublicKey *rsa.PublicKey
}

func NewFoodController(foodService service.IFoodService, rSAPublicKey *rsa.PublicKey) IFoodController {
	return &FoodController{
		FoodService:  foodService,
		RSAPublicKey: rSAPublicKey,
	}
}

func (controller *FoodController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	foodResponses := controller.FoodService.FindAll(request.Context(), limit, offset)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   foodResponses,
	}

	helper.WriteToResponseBody(writer, http.StatusAccepted, webResponse)
}

func (controller *FoodController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	IDFood := params.ByName("IDFood")
	id, err := strconv.ParseUint(IDFood, 10, 32)
	helper.PanicIfError(err)

	foodResponse := controller.FoodService.FindById(request.Context(), uint(id))

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   foodResponse,
	}

	helper.WriteToResponseBody(writer, http.StatusAccepted, webResponse)
}

func (controller *FoodController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	foodCreateRequest := model.FoodCreateRequest{}
	helper.ReadFromRequestBody(request, &foodCreateRequest)

	username := helper.GetUsername(request, controller.RSAPublicKey)
	foodCreateRequest.CreatedBy = username
	foodCreateRequest.UpdatedBy = username

	foodResponse := controller.FoodService.Create(request.Context(), foodCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   foodResponse,
	}

	helper.WriteToResponseBody(writer, http.StatusAccepted, webResponse)
}

func (controller *FoodController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	foodUpdateRequest := model.FoodUpdateRequest{}
	helper.ReadFromRequestBody(request, &foodUpdateRequest)

	IDFood := params.ByName("IDFood")
	id, err := strconv.ParseUint(IDFood, 10, 32)
	helper.PanicIfError(err)

	username := helper.GetUsername(request, controller.RSAPublicKey)
	foodUpdateRequest.IDFood = uint(id)
	foodUpdateRequest.UpdatedBy = username

	foodResponse := controller.FoodService.Update(request.Context(), foodUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   foodResponse,
	}

	helper.WriteToResponseBody(writer, http.StatusAccepted, webResponse)
}

func (controller *FoodController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	IDFood := params.ByName("IDFood")
	id, err := strconv.ParseUint(IDFood, 10, 32)
	helper.PanicIfError(err)

	foodDeleteRequest := model.FoodDeleteRequest{
		IDFood:    uint(id),
		DeletedBy: helper.GetUsername(request, controller.RSAPublicKey),
	}

	controller.FoodService.Delete(request.Context(), foodDeleteRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, http.StatusAccepted, webResponse)
}
