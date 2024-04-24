package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/error/entity"
	entityCate "github.com/textures1245/BlogDuaaeeg-backend/model/category/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/model/post/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/model/utils"
)

type PostRepo struct {
	Db *db.PrismaClient
}

func NewPostRepository(db *db.PrismaClient) entity.PostRepository {
	return &PostRepo{
		Db: db,
	}
}

func (postRepo *PostRepo) CreatePost(cateResDat *entityCate.PostCategoryResDat, tagResDat *entityCate.PostTagResDat, req *entity.PostReqDat) (*db.PostModel, error) {
	ctx := context.Background()

	if err := utils.SchemaValidator(req); err != nil {
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
		db.Post.Category.Link(
			db.PostCategory.ID.Equals(cateResDat.ID),
		),
		db.Post.Tags.Link(
			db.PostTag.ID.Equals(tagResDat.ID),
		),
	).With(
		db.Post.Category.Fetch(),
		db.Post.Tags.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}

	}
	return post, nil
}

func (postRepo *PostRepo) UpdatePostByUUID(cateResDat *entityCate.PostCategoryResDat, uuid string, req *entity.PostReqDat) (*db.PostModel, error) {
	ctx := context.Background()

	if err := utils.SchemaValidator(req); err != nil {
		return nil, err
	}
	srcTypes := []string{"MARKDOWN_URL", "CONTENT", "MARKDOWN_FILE"}

	if !utils.Contains(srcTypes, req.SrcType) {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("InvalidType, on src_type field, only allowed MARKDOWN_URL, CONTENT, MARKDOWN_FILE"),
		}
	}

	post, err := postRepo.Db.Post.FindUnique(
		db.Post.UUID.Equals(uuid),
	).With(
		db.Post.Tags.Fetch(),
		db.Post.Category.Fetch(),
	).Update(
		db.Post.Title.Set(req.Title),
		db.Post.Source.Set(req.Content),
		db.Post.SrcType.Set(db.SrcType(req.SrcType)),
		db.Post.Published.Set(req.Published),
		db.Post.Category.Link(
			db.PostCategory.ID.Equals(cateResDat.ID),
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

func (postRepo *PostRepo) UpdatePostToPublisher(userUuid string, postUuid string) error {
	ctx := context.Background()

	pp, err := postRepo.Db.PublicationPost.CreateOne(
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

	fmt.Println(postUuid, pp.ID)

	if _, err := postRepo.Db.Post.FindUnique(
		db.Post.UUID.Equals(postUuid),
	).Update(
		db.Post.PublishPostUUID.Set(pp.UUID),
		db.Post.PublicationPost.Link(
			db.PublicationPost.UUID.Equals(pp.UUID),
		),
	).Exec(ctx); err != nil {
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
	).With(
		db.Post.Category.Fetch(),
		db.Post.Tags.Fetch(),
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
		db.PublicationPost.Post.Fetch().With(
			db.Post.Category.Fetch(),
			db.Post.Tags.Fetch(),
		),
	).Skip(render * opts.Page).Take(render).Exec(ctx)
	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return posts, nil
}

func (postRepo *PostRepo) FetchPostByUserUUID(userUuid string) ([]db.PostModel, error) {
	ctx := context.Background()

	posts, err := postRepo.Db.Post.FindMany(
		db.Post.UserUUID.Equals(userUuid),
	).With(
		db.Post.Category.Fetch(),
		db.Post.Tags.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return posts, nil
}
