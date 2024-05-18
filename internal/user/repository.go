package user

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/entities"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user/dtos"
)

type UsersRepository interface {
	FindUserAsPassport(email string) (*entities.UsersPassport, error)
	GetUserByUUID(userUuid string) (*db.UserModel, error)
	GetUsers() ([]db.UserModel, error)
	CreateUser(req *entities.UsersCredentials) (*entities.UsersPassport, error)
	UpdateProfile(userUuid string, req *dtos.UserProfileDataRequest) (*db.UserProfileModel, error)
	DeleteUserByUuid(userUuid string) error
}
