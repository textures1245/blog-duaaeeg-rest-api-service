package entity

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/model/auth/entity"
)

type UsersRepository interface {
	FindUserAsPassport(email string) (*entity.UsersPassport, error)
	CreateUser(req *entity.UsersCredentials) (*entity.UsersPassport, error)
	UpdateProfile(req *UserProfileDataRequest) (*db.UserProfileModel, error)
}

type UserProfileDataRequest struct {
	UserUUID       string `json:"user_uuid" db:"user_uuid" form:"user_uuid" binding:"required" validate:"required, uuid"`
	FirstName      string `json:"first_name" db:"first_name" form:"first_name" binding:"required" validate:"required"`
	LastName       string `json:"last_name" db:"last_name" form:"last_name" binding:"required" validate:"required"`
	Bio            string `json:"bio" db:"bio" form:"bio" binding:"required" validate:"required"`
	ProfilePicture string `json:"profile_pic" db:"profile_pic" form:"profile_pic" binding:"required" validate:"required, url"`
}

type UserService interface {
	OnUpdateUserProfile(req *UserProfileDataRequest) (*UserProfileRes, error)
}

type UserProfileRes struct {
	UUID           string `json:"uuid"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profile_pic"`
}
