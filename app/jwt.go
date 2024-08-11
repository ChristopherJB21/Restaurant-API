package app

import (
	"crypto/rsa"
	"os"
	"restaurant/helper"

	"github.com/golang-jwt/jwt/v5"
)

func NewRSAPublicKey() *rsa.PublicKey{
	publicKey, err := os.ReadFile("publicKey")
	helper.PanicIfError(err)

	PublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	helper.PanicIfError(err)

	return PublicKey
}