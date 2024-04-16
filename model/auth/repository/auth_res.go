package repository

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/model/auth/entity"
)

type authRepo struct {
	Db *db.PrismaClient
}

func NewAuthRepository(db *db.PrismaClient) entity.AuthRepository {
	return &authRepo{
		Db: db,
	}
}

func (r *authRepo) SignUsersAccessToken(req *entity.UsersPassport) (string, error) {
	claims := entity.UsersClaims{
		Uuid:  req.Uuid,
		Email: req.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	mySigningKey := os.Getenv("JWT_SECRET_TOKEN")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return ss, nil
}
