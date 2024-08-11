package repository

import (
	"context"
	"errors"
	"restaurant/helper"
	model "restaurant/model/food"

	"gorm.io/gorm"
)

type IFoodRepository interface {
	FindAll(ctx context.Context, limit int, offset int) ([]model.Food, error)
	FindById(ctx context.Context, IDFood uint) (model.Food, error)
	Create(ctx context.Context, food model.Food) model.Food
	Update(ctx context.Context, food model.Food) model.Food
	Delete(ctx context.Context, food model.Food)
}

type FoodRepository struct {
	DB *gorm.DB
}

func NewFoodRepository(DB *gorm.DB) IFoodRepository {
	return &FoodRepository{
		DB: DB,
	}
}

func (repository *FoodRepository) FindAll(ctx context.Context, limit int, offset int) ([]model.Food, error) {
	var foods []model.Food

	result := repository.DB.Joins("Cuisine").Limit(limit).Offset(offset).Find(&foods)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || len(foods) < 1 {
		return foods, errors.New("foods are not found")
	}

	helper.PanicIfError(result.Error)

	return foods, nil
}

func (repository *FoodRepository) FindById(ctx context.Context, IDFood uint) (model.Food, error) {
	var food model.Food

	result := repository.DB.Joins("Cuisine").First(&food, IDFood)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return food, errors.New("food is not found")
	}

	helper.PanicIfError(result.Error)

	return food, nil
}

func (repository *FoodRepository) Create(ctx context.Context, food model.Food) model.Food {
	result := repository.DB.Create(&food)

	helper.PanicIfError(result.Error)

	return food
}

func (repository *FoodRepository) Update(ctx context.Context, food model.Food) model.Food {
	result := repository.DB.Save(&food)

	helper.PanicIfError(result.Error)

	return food
}

func (repository *FoodRepository) Delete(ctx context.Context, food model.Food) {
	result := repository.DB.Delete(&food)

	helper.PanicIfError(result.Error)
}
