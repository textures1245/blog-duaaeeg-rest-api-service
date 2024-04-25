package repository

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/textures1245/BlogDuaaeeg-backend/api/db"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/api/error/entity"
	_authEntity "github.com/textures1245/BlogDuaaeeg-backend/api/model/auth/entity"
	_userEntity "github.com/textures1245/BlogDuaaeeg-backend/api/model/user/entity"
)

type UserRepo struct {
	Db *db.PrismaClient
}

func NewUserRepository(db *db.PrismaClient) _userEntity.UsersRepository {
	return &UserRepo{
		Db: db,
	}
}

func (u *UserRepo) GetUserByUUID(usrUuid string) (*db.UserModel, error) {
	ctx := context.Background()
	user, err := u.Db.User.FindUnique(db.User.UUID.Equals(usrUuid)).With(
		db.User.UserProfile.Fetch(),
		db.User.UserFollowee.Fetch(),
		db.User.UserFollower.Fetch(),
	).Exec(ctx)
	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return user, nil
}

func (u *UserRepo) FindUserAsPassport(email string) (*_authEntity.UsersPassport, error) {
	ctx := context.Background()
	user, err := u.Db.User.FindUnique(db.User.Email.Equals(email)).Exec(ctx)
	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusNotFound,
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
		Uuid:      user.UUID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	return res, nil
}

func (u *UserRepo) UpdateProfile(userUuid string, req *_userEntity.UserProfileDataRequest) (*db.UserProfileModel, error) {
	ctx := context.Background()

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(req); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			return nil, &_errEntity.CError{
				StatusCode: http.StatusBadRequest,
				Err:        errors,
			}
		}
	}

	updateUsrProfile, err := u.createOrUpdateUserProfile(&ctx, userUuid, req)
	if err != nil {
		return nil, err
	}

	return updateUsrProfile, nil
}

func (u *UserRepo) createOrUpdateUserProfile(ctx *context.Context, userUuid string, req *_userEntity.UserProfileDataRequest) (*db.UserProfileModel, error) {
	userProfile, err := u.Db.UserProfile.UpsertOne(
		db.UserProfile.UserUUID.Equals(userUuid),
	).Create(
		db.UserProfile.FirstName.Set(req.FirstName),
		db.UserProfile.LastName.Set(req.LastName),
		db.UserProfile.Bio.Set(req.Bio),
		db.UserProfile.ProfilePicture.Set(req.ProfilePicture),
		db.UserProfile.User.Link(
			db.User.UUID.Equals(userUuid),
		),
	).Update(
		db.UserProfile.FirstName.Set(req.FirstName),
		db.UserProfile.LastName.Set(req.LastName),
		db.UserProfile.Bio.Set(req.Bio),
		db.UserProfile.ProfilePicture.Set(req.ProfilePicture),
	).Exec(*ctx)
	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return userProfile, nil
}
