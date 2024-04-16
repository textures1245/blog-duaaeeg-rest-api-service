package entity

import (
	"github.com/golang-jwt/jwt/v4"
)

type AuthRepository interface {
	SignUsersAccessToken(req *UsersPassport) (string, error)
}

type AuthService interface {
	Login(req *UsersCredentials) (*UsersLoginRes, error)
	Register(req *UsersCredentials) (*UsersLoginRes, error)
}

type UsersCredentials struct {
	Email    string `json:"email" db:"email" form:"email" binding:"required"`
	Password string `json:"password" db:"password" form:"password" binding:"required"`
}

type UsersPassport struct {
	Uuid     string `json:"uuid" db:"uuid"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UsersClaims struct {
	Uuid  string `json:"user_uuid"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type UsersLoginRes struct {
	AccessToken string `json:"access_token"`
}
