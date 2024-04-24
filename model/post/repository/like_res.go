package repository

import (
	"context"
	"net/http"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/model/post/entity"
)

type LikeRepo struct {
	db *db.PrismaClient
}

func NewLikeRepo(db *db.PrismaClient) entity.LikeRepository {
	return &LikeRepo{db}
}

func (l *LikeRepo) CreateLike(pUuid string, req *entity.LikeReqDat) (*db.LikeModel, error) {
	ctx := context.Background()

	like, err := l.db.Like.CreateOne(
		db.Like.User.Link(
			db.User.UUID.Equals(req.UserUuid),
		),
		db.Like.Post.Link(
			db.Post.UUID.Equals(pUuid),
		),
	).Exec(ctx)
	if err != nil {
		return nil, &_errEntity.CError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}
	return like, nil
}

func (l *LikeRepo) DeleteLikeByUUID(pUuid string, usrUuid string) error {
	ctx := context.Background()

	// TODO: Implement Custom Raw Query form delete like (DONE)

	_, err := l.db.Prisma.ExecuteRaw(`DELETE FROM "Like" WHERE "postUuid" = $1 AND "userUuid" = $2`, pUuid, usrUuid).Exec(ctx)
	if err != nil {
		return &_errEntity.CError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}
