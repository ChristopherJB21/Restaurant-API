package service

import (
	"context"
	"database/sql"
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

	// foods := []model.FoodResponse{}

	// for i := 1; i <= 5; i++ {
	// 	food := model.FoodResponse{}
	// 	food.IDFood = uint(i)
	// 	food.FoodName = "Hello World"
	// 	food.IDCuisine = uint(i + 1)
	// 	food.CuisineName = "Hai World"
	// 	foods = append(foods, food)
	// }

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
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	food := model.Food{}
	food.FoodName = request.FoodName
	food.IDCuisine = request.IDCuisine
	food.CreatedBy = request.CreatedBy
	food.UpdatedBy = request.UpdatedBy

	food = service.FoodRepository.Create(ctx, food)

	food, err = service.FoodRepository.FindById(ctx, food.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToFoodResponse(food)
}

func (service *FoodService) Update(ctx context.Context, request model.FoodUpdateRequest) model.FoodResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	food, err := service.FoodRepository.FindById(ctx, request.IDFood)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	food.FoodName = request.FoodName
	food.IDCuisine = request.IDCuisine
	food.UpdatedBy = request.UpdatedBy

	food = service.FoodRepository.Update(ctx, food)

	food, err = service.FoodRepository.FindById(ctx, request.IDFood)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToFoodResponse(food)
}

func (service *FoodService) Delete(ctx context.Context, request model.FoodDeleteRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	food, err := service.FoodRepository.FindById(ctx, request.IDFood)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	food.DeletedBy = sql.NullString{String: request.DeletedBy, Valid: true}

	service.FoodRepository.Delete(ctx, food)
}
