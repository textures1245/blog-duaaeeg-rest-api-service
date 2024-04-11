package entity

import (
	"github.com/textures1245/BlogDuaaeeg-backend/model/auth/entity"
)

type UsersRepository interface {
	FindUserAsPassport(email string) (*entity.UsersPassport, error)
	CreateUser(req *entity.UsersCredentials) (*entity.UsersPassport, error)
}
