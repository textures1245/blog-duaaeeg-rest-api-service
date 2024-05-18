package usecase

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/entities"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/utils"

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

func (u *authUse) Login(req *entities.UsersCredentials, hashMethod ...string) (*dtos.UsersLoginRes, error) {

	user, err := u.UsersRepo.FindUserAsPassport(req.Email)
	if err != nil {
		return nil, err
	}

	userPwdDecode, errOnDecode := base64.StdEncoding.DecodeString(user.Password)
	if errOnDecode != nil {
		return nil, &entity.CError{
			Err:        err,
			StatusCode: 500,
		}
	}

	if hashMethod[0] == "AES" {
		key32, error := utils.ToByte32([]byte(os.Getenv("AES_KEY")))
		if error != nil {
			return nil, error
		}

		if err := utils.AESHashCompared([]byte(req.Password), userPwdDecode, key32); err != nil {
			return nil, err
		}
	} else {
		if err := bcrypt.CompareHashAndPassword(userPwdDecode, []byte(req.Password)); err != nil {
			fmt.Println(err.Error())
			return nil, errors.New("error, password is invalid")
		}
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

func (u *authUse) Register(req *entities.UsersCredentials, hashMethod ...string) (*dtos.UsersLoginRes, error) {

	var hashedPassword []byte
	if hashMethod[0] == "AES" {
		key32, error := utils.ToByte32([]byte(os.Getenv("AES_KEY")))
		if error != nil {
			return nil, error
		}

		hashedPasswordFromAES, err := utils.AESEncryption([]byte(req.Password), key32)
		if err != nil {
			return nil, err
		}
		hashedPassword = hashedPasswordFromAES
	} else {
		hashedPasswordFromBcrypt, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		hashedPassword = hashedPasswordFromBcrypt

	}

	cred := entities.UsersCredentials{
		Email:    req.Email,
		Password: base64.StdEncoding.EncodeToString(hashedPassword),
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
