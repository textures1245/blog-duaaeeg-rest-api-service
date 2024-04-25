package repository

import (
	"context"
	"net/http"

	"github.com/textures1245/BlogDuaaeeg-backend/api/db"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/api/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/api/model/post/entity"
)

type CommRepo struct {
	db *db.PrismaClient
}

func NewCommRepo(db *db.PrismaClient) entity.CommentRepository {
	return &CommRepo{db}
}

func (c *CommRepo) CreateComment(pUuid string, req *entity.CommentReqDat) (*db.CommentModel, error) {
	ctx := context.Background()

	comm, err := c.db.Comment.CreateOne(
		db.Comment.Content.Set(req.Comment),
		db.Comment.User.Link(
			db.User.UUID.Equals(req.UserUuid),
		),
		db.Comment.Post.Link(
			db.Post.UUID.Equals(pUuid),
		),
	).Exec(ctx)
	if err != nil {
		return nil, &_errEntity.CError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}
	return comm, nil
}

func (c *CommRepo) UpdateCommentByUUID(cUuid string, req *entity.CommentReqDat) (*db.CommentModel, error) {
	ctx := context.Background()

	comm, err := c.db.Comment.FindUnique(
		db.Comment.UUID.Equals(cUuid),
	).Update(
		db.Comment.Content.Set(req.Comment),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return comm, nil
}

func (c *CommRepo) DeleteCommentByUUID(cUuid string) error {
	ctx := context.Background()

	_, err := c.db.Comment.FindUnique(
		db.Comment.UUID.Equals(cUuid),
	).Delete().Exec(ctx)
	if err != nil {
		return &_errEntity.CError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}
	return nil
}
