package repository

import (
	"context"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
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

	user, err := u.Db.User.CreateOne(
		db.User.Email.Set(req.Email),
		db.User.Password.Set(req.Password),
	).Exec(ctx)
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
