package web

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type SSOClaims struct {
	IDUser   uuid.UUID `json:"iduser"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}
