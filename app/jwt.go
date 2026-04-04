package app

import (
	"crypto/rsa"
	"os"
	"restaurant/helper"

	"github.com/golang-jwt/jwt/v5"
)

func NewRSAPublicKey() *rsa.PublicKey {
	publicKey, err := os.ReadFile("publicKey")
	helper.PanicIfError(err)

	PublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	helper.PanicIfError(err)

	return PublicKey
}

func NewRSAPrivateKey() *rsa.PrivateKey {
	privateKey, err := os.ReadFile("privateKey")
	helper.PanicIfError(err)

	PrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	helper.PanicIfError(err)

	return PrivateKey
}
