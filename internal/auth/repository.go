package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/entities"
)

type AuthRepository interface {
	SignUsersAccessToken(req *entities.UsersPassport) (string, *jwt.NumericDate, error)
}
