package app

import (
	"crypto/rsa"
	"restaurant/controller"
	"restaurant/exception"
	"restaurant/repository"
	"restaurant/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewRouter(DB *gorm.DB, validate *validator.Validate, rSAPublicKey *rsa.PublicKey, rSAPrivateKey *rsa.PrivateKey, redis *redis.Client) *httprouter.Router {
	router := httprouter.New()

	router.PanicHandler = exception.ErrorHandler

	NewCuisineRouter(router, DB, validate, rSAPublicKey, redis)
	NewFoodRouter(router, DB, validate, rSAPublicKey)
	NewUserRouter(router, DB, validate, rSAPrivateKey)

	return router
}

func NewCuisineRouter(router *httprouter.Router, DB *gorm.DB, validate *validator.Validate, rSAPublicKey *rsa.PublicKey, redis *redis.Client) {
	cuisineRepository := repository.NewCuisineRepository(DB)
	cuisineService := service.NewCuisineService(cuisineRepository, validate, redis)
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

func NewUserRouter(router *httprouter.Router, DB *gorm.DB, validate *validator.Validate, rSAPrivateKey *rsa.PrivateKey) {
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepository, validate, rSAPrivateKey)
	userController := controller.NewUserController(userService)

	router.GET("/api/users", userController.FindAll)
	router.GET("/api/user/:IDUser", userController.FindById)
	router.POST("/api/user/login", userController.Login)
	router.POST("/api/user", userController.Create)
	router.PUT("/api/userpassword/:IDUser", userController.UpdatePassword)
	router.PUT("/api/user/:IDUser", userController.Update)
	router.DELETE("/api/user/:IDUser", userController.Delete)
}
