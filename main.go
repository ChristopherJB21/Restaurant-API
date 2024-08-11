package main

import (
	"restaurant/app"
	"restaurant/helper"
	"restaurant/middleware"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func main() {
	app.NewViper()

	DB := app.NewDB()

	validate := validator.New()
	
	// redis := app.NewRedis()

	rSAPublicKey := app.NewRSAPublicKey()
	
	router := app.NewRouter(DB, validate, rSAPublicKey)

	server := http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: middleware.NewMiddleware(router, validate, rSAPublicKey),
	}

	log.Println(viper.GetString("appName") + " Application Start")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
