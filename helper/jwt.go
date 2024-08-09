package helper

import (
	"errors"
	"net/http"
	"os"
	model "restaurant/model/web"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetToken(request *http.Request) (*jwt.Token, error) {
	BearerKey := request.Header.Get("Authorization")
	tokenString := strings.TrimSpace(strings.Replace(BearerKey, "Bearer", "", 1))

	publicKey, err := os.ReadFile("publicKey")
	PanicIfError(err)

	PublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	PanicIfError(err)

	token, err := jwt.ParseWithClaims(tokenString, &model.SSOClaims{}, func(token *jwt.Token) (interface{}, error) {
		return PublicKey, nil
	}, jwt.WithLeeway(5*time.Second))

	return token, err
}

func VerifyToken(request *http.Request) (err error) {
	token, err := GetToken(request)

	if err != nil {
		return err
	} else if !token.Valid {
		return errors.New("token invalid")
	}

	return nil
}

func GetUsername(request *http.Request) (username string) {
	token, _ := GetToken(request)

	claims := token.Claims.(*model.SSOClaims)

	return claims.Username
}
