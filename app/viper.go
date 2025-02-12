package app

import (
	"restaurant/helper"

	"github.com/spf13/viper"
)

func NewViper() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("api.config")

	err := viper.ReadInConfig()
	helper.PanicIfError(err)
}