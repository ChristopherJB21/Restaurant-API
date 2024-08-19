package app

import (
	"restaurant/controller"
	"restaurant/exception"
	"restaurant/repository"
	"restaurant/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func NewRouter(DB *gorm.DB, validate *validator.Validate) *httprouter.Router {
	// Prepare HTTP Router
	router := httprouter.New()

	// Set Panic Handler
	router.PanicHandler = exception.ErrorHandler

	NewCuisineRouter(router, DB, validate)
	NewFoodRouter(router, DB, validate)

	return router
}

func NewCuisineRouter(router *httprouter.Router, DB *gorm.DB, validate *validator.Validate) {
	cuisineRepository := repository.NewCuisineRepository(DB)
	cuisineService := service.NewCuisineService(cuisineRepository, validate)
	cuisineController := controller.NewCuisineController(cuisineService)

	router.GET("/api/cuisines", cuisineController.FindAll)
	router.GET("/api/cuisine/:IDCuisine", cuisineController.FindById)
	router.POST("/api/cuisine", cuisineController.Create)
	router.PUT("/api/cuisine/:IDCuisine", cuisineController.Update)
	router.DELETE("/api/cuisine/:IDCuisine", cuisineController.Delete)
}

func NewFoodRouter(router *httprouter.Router, DB *gorm.DB, validate *validator.Validate) {
	foodRepository := repository.NewFoodRepository(DB)
	foodService := service.NewFoodService(foodRepository, validate)
	foodController := controller.NewFoodController(foodService)
		
	router.GET("/api/foods", foodController.FindAll)
	router.GET("/api/food/:IDFood", foodController.FindById)
	router.POST("/api/food", foodController.Create)
	router.PUT("/api/food/:IDFood", foodController.Update)
	router.DELETE("/api/food/:IDFood", foodController.Delete)
}