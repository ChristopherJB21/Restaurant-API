package app

import (
	"crypto/rsa"
	"restaurant/controller"
	"restaurant/exception"
	"restaurant/repository"
	"restaurant/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func NewRouter(DB *gorm.DB, validate *validator.Validate, rSAPublicKey *rsa.PublicKey) *httprouter.Router {
	router := httprouter.New()

	router.PanicHandler = exception.ErrorHandler

	NewCuisineRouter(router, DB, validate, rSAPublicKey)
	NewFoodRouter(router, DB, validate, rSAPublicKey)

	return router
}

func NewCuisineRouter(router *httprouter.Router, DB *gorm.DB, validate *validator.Validate, rSAPublicKey *rsa.PublicKey) {
	cuisineRepository := repository.NewCuisineRepository(DB)
	cuisineService := service.NewCuisineService(cuisineRepository, validate)
	cuisineController := controller.NewCuisineController(cuisineService, rSAPublicKey)

	router.GET("/api/cuisines", cuisineController.FindAll)
	router.GET("/api/cuisine/:IDCuisine", cuisineController.FindById)
	router.POST("/api/cuisine", cuisineController.Create)
	router.PUT("/api/cuisine/:IDCuisine", cuisineController.Update)
	router.DELETE("/api/cuisine/:IDCuisine", cuisineController.Delete)
}

func NewFoodRouter(router *httprouter.Router, DB *gorm.DB, validate *validator.Validate, rSAPublicKey *rsa.PublicKey) {
	foodRepository := repository.NewFoodRepository(DB)
	foodService := service.NewFoodService(foodRepository, validate)
	foodController := controller.NewFoodController(foodService, rSAPublicKey)
		
	router.GET("/api/foods", foodController.FindAll)
	router.GET("/api/food/:IDFood", foodController.FindById)
	router.POST("/api/food", foodController.Create)
	router.PUT("/api/food/:IDFood", foodController.Update)
	router.DELETE("/api/food/:IDFood", foodController.Delete)
}