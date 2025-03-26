package helper

import (
	"crypto/rsa"
	"errors"
	"net/http"
	model "restaurant/model/web"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetToken(request *http.Request, rSAPublicKey *rsa.PublicKey) (*jwt.Token, error) {
	BearerKey := request.Header.Get("Authorization")
	tokenString := strings.TrimSpace(strings.Replace(BearerKey, "Bearer", "", 1))

	token, err := jwt.ParseWithClaims(tokenString, &model.SSOClaims{}, func(token *jwt.Token) (interface{}, error) {
		return rSAPublicKey, nil
	}, jwt.WithLeeway(5*time.Second))

	return token, err
}

func VerifyToken(request *http.Request, rSAPublicKey *rsa.PublicKey) (err error) {
	token, err := GetToken(request, rSAPublicKey)

	if err != nil {
		return err
	} else if !token.Valid {
		return errors.New("token invalid")
	}

	return nil
}

func GetUsername(request *http.Request, rSAPublicKey *rsa.PublicKey) (username string) {
	token, _ := GetToken(request, rSAPublicKey)

	claims := token.Claims.(*model.SSOClaims)

	return claims.Username
}
