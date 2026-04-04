package helper

import (
	"crypto/rsa"
	"errors"
	"net/http"
	model_user "restaurant/model/user"
	"restaurant/model/web"
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

func GenerateToken(user model_user.User, rSAPrivateKey *rsa.PrivateKey) (string, error) {
	claims := web.SSOClaims{
		IDUser:   user.ID,
		Username: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "myrestaurant.com",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(rSAPrivateKey)
	PanicIfError(err)

	return tokenString, nil
}
