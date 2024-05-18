package auth

import (
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/entities"
)

type AuthUsecase interface {
	Login(req *entities.UsersCredentials, hashMethod ...string) (*dtos.UsersLoginRes, error)
	Register(req *entities.UsersCredentials, hashMethod ...string) (*dtos.UsersLoginRes, error)
}
