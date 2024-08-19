package service

import (
	"context"
	"restaurant/exception"
	"restaurant/helper"
	model "restaurant/model/food"
	"restaurant/repository"

	"github.com/go-playground/validator/v10"
)

type IFoodService interface {
	FindAll(ctx context.Context, limit int, offset int) []model.FoodResponse
	FindById(ctx context.Context, IDFood uint) model.FoodResponse
	Create(ctx context.Context, request model.FoodCreateRequest) model.FoodResponse
	Update(ctx context.Context, request model.FoodUpdateRequest) model.FoodResponse
	Delete(ctx context.Context, request model.FoodDeleteRequest)
}

type FoodService struct {
	FoodRepository repository.IFoodRepository
	Validate       *validator.Validate
}

func NewFoodService(foodRepository repository.IFoodRepository, validate *validator.Validate) IFoodService {
	return &FoodService{
		FoodRepository: foodRepository,
		Validate:       validate,
	}
}

func (service *FoodService) FindAll(ctx context.Context, limit int, offset int) []model.FoodResponse {
	foods, err := service.FoodRepository.FindAll(ctx, limit, offset)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToFoodResponses(foods)
}

func (service *FoodService) FindById(ctx context.Context, IDFood uint) model.FoodResponse {
	food, err := service.FoodRepository.FindById(ctx, IDFood)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToFoodResponse(food)
}

func (service *FoodService) Create(ctx context.Context, request model.FoodCreateRequest) model.FoodResponse {
	// Validate Create Request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	food := model.Food{}
	food.FoodName = request.FoodName
	food.IDCuisine = request.IDCuisine

	food = service.FoodRepository.Create(ctx, food)

	food, err = service.FoodRepository.FindById(ctx, food.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToFoodResponse(food)
}

func (service *FoodService) Update(ctx context.Context, request model.FoodUpdateRequest) model.FoodResponse {
	// Validate Update Request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	food, err := service.FoodRepository.FindById(ctx, request.IDFood)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	food.FoodName = request.FoodName
	food.IDCuisine = request.IDCuisine

	food = service.FoodRepository.Update(ctx, food)

	food, err = service.FoodRepository.FindById(ctx, request.IDFood)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToFoodResponse(food)
}

func (service *FoodService) Delete(ctx context.Context, request model.FoodDeleteRequest) {
	// Validate Delete Request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	food, err := service.FoodRepository.FindById(ctx, request.IDFood)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.FoodRepository.Delete(ctx, food)
}
