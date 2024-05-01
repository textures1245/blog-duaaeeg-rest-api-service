package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

type UsersClaims struct {
	Uuid  string `json:"user_uuid"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
