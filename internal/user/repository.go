package user

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type UsersRepository interface {
	FindUserAsPassport(email string) (*entity.UsersPassport, error)
	GetUserByUUID(userUuid string) (*db.UserModel, error)
	CreateUser(req *entity.UsersCredentials) (*entity.UsersPassport, error)
	UpdateProfile(userUuid string, req *UserProfileDataRequest) (*db.UserProfileModel, error)
}
