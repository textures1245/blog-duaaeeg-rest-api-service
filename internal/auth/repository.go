package auth

import (
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/entities"
)

type AuthRepository interface {
	SignUsersAccessToken(req *entities.UsersPassport) (string, error)
}
