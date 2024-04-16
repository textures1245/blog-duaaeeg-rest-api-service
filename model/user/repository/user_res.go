package repository

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/error/entity"
	_authEntity "github.com/textures1245/BlogDuaaeeg-backend/model/auth/entity"
	_userEntity "github.com/textures1245/BlogDuaaeeg-backend/model/user/entity"
)

type UserRepo struct {
	Db *db.PrismaClient
}

func NewUserRepository(db *db.PrismaClient) _userEntity.UsersRepository {
	return &UserRepo{
		Db: db,
	}
}

func (u *UserRepo) FindUserAsPassport(email string) (*_authEntity.UsersPassport, error) {
	ctx := context.Background()
	user, err := u.Db.User.FindUnique(db.User.Email.Equals(email)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	res := &_authEntity.UsersPassport{
		Uuid:     user.UUID,
		Email:    user.Email,
		Password: user.Password,
	}

	return res, nil
}

func (u *UserRepo) CreateUser(req *_authEntity.UsersCredentials) (*_authEntity.UsersPassport, error) {
	ctx := context.Background()

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return nil, &_errEntity.CError{
			StatusCode: http.StatusBadRequest,
			Err:        errors,
		}
	}

	user, err := u.Db.User.CreateOne(
		db.User.Email.Set(req.Email),
		db.User.Password.Set(req.Password),
	).Exec(ctx)
	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	res := &_authEntity.UsersPassport{
		Uuid:     user.UUID,
		Email:    user.Email,
		Password: user.Password,
	}

	return res, nil
}

func (u *UserRepo) UpdateProfile(req *_userEntity.UserProfileDataRequest) (*db.UserProfileModel, error) {
	ctx := context.Background()

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return nil, &_errEntity.CError{
			StatusCode: http.StatusBadRequest,
			Err:        errors,
		}
	}

	user, err := u.Db.UserProfile.CreateOne(
		db.UserProfile.FirstName.Set(req.FirstName),
		db.UserProfile.LastName.Set(req.LastName),
		db.UserProfile.Bio.Set(req.Bio),
		db.UserProfile.ProfilePicture.Set(req.ProfilePicture),

		db.UserProfile.User.Link(
			db.User.UUID.Equals(req.UserUUID),
		),
	).Exec(ctx)
	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return user, nil

}
