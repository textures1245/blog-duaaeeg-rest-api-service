package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/model/post/entity"
)

type PostRepo struct {
	Db *db.PrismaClient
}

func NewPostRepository(db *db.PrismaClient) entity.PostRepository {
	return &PostRepo{
		Db: db,
	}
}

func (postRepo *PostRepo) CreatePost(req *entity.PostReqDat) (*db.PostModel, error) {
	ctx := context.Background()

	if err := postValidator(req); err != nil {
		return nil, err
	}

	post, err := postRepo.Db.Post.CreateOne(
		db.Post.Title.Set(req.Title),
		db.Post.Source.Set(req.Content),
		db.Post.SrcType.Set(db.SrcType(req.SrcType)),
		db.Post.Published.Set(req.Published),
		db.Post.User.Link(
			db.User.UUID.Equals(req.UserUuid),
		),
	).Exec(ctx)

	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}

	}
	return post, nil
}

func (postRepo *PostRepo) UpdatePostByUUID(uuid string, req *entity.PostReqDat) (*db.PostModel, error) {
	ctx := context.Background()

	if err := postValidator(req); err != nil {
		return nil, err
	}

	post, err := postRepo.Db.Post.FindUnique(
		db.Post.UUID.Equals(uuid),
	).Update(
		db.Post.Title.Set(req.Title),
		db.Post.Source.Set(req.Content),
		db.Post.SrcType.Set(db.SrcType(req.SrcType)),
		db.Post.Published.Set(req.Published),
	).Exec(ctx)

	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return post, nil
}

func (postRepo *PostRepo) UpdatePostToPublisher(userUuid string, postUuid string) error {
	ctx := context.Background()

	_, err := postRepo.Db.PublicationPost.CreateOne(
		db.PublicationPost.User.Link(
			db.User.UUID.Equals(userUuid),
		),
		db.PublicationPost.Post.Link(
			db.Post.UUID.Equals(postUuid),
		),
	).Exec(ctx)

	if err != nil {
		return &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return nil
}

func (postRepo *PostRepo) FetchPostByUUID(uuid string) (*db.PostModel, error) {
	ctx := context.Background()

	post, err := postRepo.Db.Post.FindUnique(
		db.Post.UUID.Equals(uuid),
	).Exec(ctx)

	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return post, nil
}

func (postRepo *PostRepo) FetchPublisherPosts(opts *entity.FetchPostOptReq) ([]db.PublicationPostModel, error) {
	ctx := context.Background()

	if opts.Page < 0 {
		opts.Page = 0
	}

	render := 10

	posts, err := postRepo.Db.PublicationPost.FindMany().With(
		db.PublicationPost.Post.Fetch(),
	).Skip(render * opts.Page).Take(render).Exec(ctx)
	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return posts, nil
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func (postRepo *PostRepo) FetchPostByUserUUID(userUuid string) ([]db.PostModel, error) {
	ctx := context.Background()

	posts, err := postRepo.Db.Post.FindMany(
		db.Post.UserUUID.Equals(userUuid),
	).Exec(ctx)

	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return posts, nil
}

func postValidator(req *entity.PostReqDat) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return &_errEntity.CError{
			StatusCode: http.StatusBadRequest,
			Err:        errors,
		}
	}

	srcTypes := []string{"MARKDOWN_URL", "CONTENT", "MARKDOWN_FILE"}

	if !contains(srcTypes, req.SrcType) {
		return &_errEntity.CError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("InvalidType, on src_type field, only allowed MARKDOWN_URL, CONTENT, MARKDOWN_FILE"),
		}
	}
	return nil
}
