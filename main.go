package main

import (
	"restaurant/app"
	"restaurant/helper"
	"restaurant/middleware"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func main() {
	app.NewViper()

	DB := app.NewDB()

	validate := validator.New()
	
	// redis := app.NewRedis()

	router := app.NewRouter(DB, validate)

	server := http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: middleware.NewMiddleware(router, validate),
	}

	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	log.Println(viper.GetString("appName") + " Application Start")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
