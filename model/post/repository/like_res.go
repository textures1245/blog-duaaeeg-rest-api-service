package repository

import (
	"context"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/model/post/entity"
)

type LikeRepo struct {
	db *db.PrismaClient
}

func NewLikeRepo(db *db.PrismaClient) entity.LikeRepository {
	return &LikeRepo{db}
}

func (l *LikeRepo) CreateLike(req *entity.LikeReqDat) (*db.LikeModel, error) {
	ctx := context.Background()

	like, err := l.db.Like.CreateOne(
		db.Like.User.Link(
			db.User.UUID.Equals(req.UserUuid),
		),
		db.Like.Post.Link(
			db.Post.UUID.Equals(req.PostUuid),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (l *LikeRepo) DeleteLikeByUUID(uuid string) error {
	ctx := context.Background()

	_, err := l.db.Like.FindUnique(
		db.Like.UUID.Equals(uuid),
	).Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
