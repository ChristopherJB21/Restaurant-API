package main

import (
	"log"
	"net/http"
	"restaurant/app"
	"restaurant/helper"
	"restaurant/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func main() {
	// Prepare Configuration Management
	app.NewViper()

	// Prepare Database Connection
	DB := app.NewDB()

	// Prepare Request Validation
	validate := validator.New()
	
	// Prepare Route
	router := app.NewRouter(DB, validate)

	// Prepare Server
	server := http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: middleware.NewMiddleware(router),
	}

	// Start API
	log.Println(viper.GetString("appName") + " Application Start")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
