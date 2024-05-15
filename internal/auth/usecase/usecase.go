package usecase

import (
	"errors"
	"fmt"

	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/entities"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type authUse struct {
	AuthRepo  auth.AuthRepository
	UsersRepo user.UsersRepository
}

func NewAuthService(authRepo auth.AuthRepository, usersRepo user.UsersRepository) auth.AuthUsecase {
	return &authUse{
		AuthRepo:  authRepo,
		UsersRepo: usersRepo,
	}
}

func (u *authUse) Login(req *entities.UsersCredentials) (*dtos.UsersLoginRes, error) {

	user, err := u.UsersRepo.FindUserAsPassport(req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error, password is invalid")
	}

	token, err := u.AuthRepo.SignUsersAccessToken(user)
	if err != nil {
		return nil, err
	}
	res := &dtos.UsersLoginRes{
		AccessToken: token,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	return res, nil
}

func (u *authUse) Register(req *entities.UsersCredentials) (*dtos.UsersLoginRes, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	cred := entities.UsersCredentials{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	user, err := u.UsersRepo.CreateUser(&cred)
	if err != nil {
		return nil, err
	}

	token, err := u.AuthRepo.SignUsersAccessToken(user)
	if err != nil {
		return nil, err
	}
	res := &dtos.UsersLoginRes{
		AccessToken: token,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return res, nil

}
