package service

import (
	"context"
	"database/sql"
	"restaurant/exception"
	"restaurant/helper"
	model "restaurant/model/cuisine"
	"restaurant/repository"

	"github.com/go-playground/validator/v10"
)

type ICuisineService interface {
	Create(ctx context.Context, request model.CuisineCreateRequest) model.CuisineResponse
	Delete(ctx context.Context, request model.CuisineDeleteRequest)
	Update(ctx context.Context, request model.CuisineUpdateRequest) model.CuisineResponse
	FindAll(ctx context.Context, limit int, offset int) []model.CuisineResponse
	FindById(ctx context.Context, IDCuisine uint) model.CuisineResponse
}

type CuisineService struct {
	CuisineRepository repository.ICuisineRepository
	Validate          *validator.Validate
}

func NewCuisineService(cuisineRepository repository.ICuisineRepository, validate *validator.Validate) ICuisineService {
	return &CuisineService{
		CuisineRepository: cuisineRepository,
		Validate:          validate,
	}
}

func (service *CuisineService) Create(ctx context.Context, request model.CuisineCreateRequest) model.CuisineResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	cuisine := model.Cuisine{
		CuisineName: request.CuisineName,
		CreatedBy: request.CreatedBy,
		UpdatedBy: request.UpdatedBy,
	}

	cuisine = service.CuisineRepository.Create(ctx, cuisine)

	cuisine, err = service.CuisineRepository.FindById(ctx, cuisine.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToCuisineResponse(cuisine)
}

func (service *CuisineService) Delete(ctx context.Context, request model.CuisineDeleteRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	
	cuisine, err := service.CuisineRepository.FindById(ctx, request.IDCuisine)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	cuisine.DeletedBy = sql.NullString{String: request.DeletedBy, Valid: true}

	service.CuisineRepository.Delete(ctx, cuisine)
}

func (service *CuisineService) Update(ctx context.Context, request model.CuisineUpdateRequest) model.CuisineResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	
	cuisine, err := service.CuisineRepository.FindById(ctx, request.IDCuisine)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	cuisine.CuisineName = request.CuisineName
	cuisine.UpdatedBy = request.UpdatedBy

	cuisine = service.CuisineRepository.Update(ctx, cuisine)

	cuisine, err = service.CuisineRepository.FindById(ctx, request.IDCuisine)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToCuisineResponse(cuisine)
}

func (service *CuisineService) FindAll(ctx context.Context, limit int, offset int) []model.CuisineResponse {
	cuisines, err := service.CuisineRepository.FindAll(ctx, limit, offset)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToCuisineResponses(cuisines)
}

func (service *CuisineService) FindById(ctx context.Context, IDCuisine uint) model.CuisineResponse {
	cuisine, err := service.CuisineRepository.FindById(ctx, IDCuisine)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToCuisineResponse(cuisine)
}
