package comment

import (
	"context"
	"net/http"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	_postInter "github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive/dtos"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
)

type CommRepo struct {
	db *db.PrismaClient
}

func NewCommRepo(db *db.PrismaClient) _postInter.CommentRepository {
	return &CommRepo{db}
}

func (c *CommRepo) CreateComment(pUuid string, req *dtos.CommentReqDat) (*db.CommentModel, error) {
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

func (c *CommRepo) UpdateCommentByUUID(cUuid string, req *dtos.CommentReqDat) (*db.CommentModel, error) {
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
