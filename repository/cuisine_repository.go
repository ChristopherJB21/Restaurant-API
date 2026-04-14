package repository

import (
	"context"
	"errors"
	"restaurant/helper"
	model "restaurant/model/cuisine"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ICuisineRepository interface {
	Create(ctx context.Context, cuisine model.Cuisine) model.Cuisine
	Delete(ctx context.Context, cuisine model.Cuisine)
	Update(ctx context.Context, cuisine model.Cuisine) model.Cuisine
	FindAll(ctx context.Context) ([]model.Cuisine, error)
	FindById(ctx context.Context, IDCuisine uuid.UUID) (model.Cuisine, error)
}

type CuisineRepository struct {
	DB *gorm.DB
}

func NewCuisineRepository(DB *gorm.DB) ICuisineRepository {
	return &CuisineRepository{
		DB: DB,
	}
}

func (repository *CuisineRepository) Create(ctx context.Context, cuisine model.Cuisine) model.Cuisine {
	result := repository.DB.Create(&cuisine)

	helper.PanicIfError(result.Error)

	return cuisine
}

func (repository *CuisineRepository) Delete(ctx context.Context, cuisine model.Cuisine) {
	result := repository.DB.Model(&cuisine).UpdateColumns(model.Cuisine{DeletedBy: cuisine.DeletedBy}).Delete(&cuisine)

	helper.PanicIfError(result.Error)
}

func (repository *CuisineRepository) Update(ctx context.Context, cuisine model.Cuisine) model.Cuisine {
	result := repository.DB.Save(&cuisine)

	helper.PanicIfError(result.Error)

	return cuisine
}

func (repository *CuisineRepository) FindAll(ctx context.Context) ([]model.Cuisine, error) {
	var cuisines []model.Cuisine

	result := repository.DB.Find(&cuisines)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || len(cuisines) < 1 {
		return cuisines, errors.New("cuisines are not found")
	}

	helper.PanicIfError(result.Error)

	return cuisines, nil
}

func (repository *CuisineRepository) FindById(ctx context.Context, IDCuisine uuid.UUID) (model.Cuisine, error) {
	var cuisine model.Cuisine

	result := repository.DB.First(&cuisine, IDCuisine)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return cuisine, errors.New("cuisine is not found")
	}

	helper.PanicIfError(result.Error)

	return cuisine, nil
}
