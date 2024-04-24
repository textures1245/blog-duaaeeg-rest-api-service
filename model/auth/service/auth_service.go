package service

import (
	"errors"
	"fmt"

	authEntity "github.com/textures1245/BlogDuaaeeg-backend/model/auth/entity"
	userEntity "github.com/textures1245/BlogDuaaeeg-backend/model/user/entity"

	"golang.org/x/crypto/bcrypt"
)

type authUse struct {
	AuthRepo  authEntity.AuthRepository
	UsersRepo userEntity.UsersRepository
}

func NewAuthService(authRepo authEntity.AuthRepository, usersRepo userEntity.UsersRepository) authEntity.AuthService {
	return &authUse{
		AuthRepo:  authRepo,
		UsersRepo: usersRepo,
	}
}

func (u *authUse) Login(req *authEntity.UsersCredentials) (*authEntity.UsersLoginRes, error) {

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
	res := &authEntity.UsersLoginRes{
		AccessToken: token,
	}
	return res, nil
}

func (u *authUse) Register(req *authEntity.UsersCredentials) (*authEntity.UsersLoginRes, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	cred := authEntity.UsersCredentials{
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
	res := &authEntity.UsersLoginRes{
		AccessToken: token,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return res, nil

}
