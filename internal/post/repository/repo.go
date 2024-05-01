package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	entityCate "github.com/textures1245/BlogDuaaeeg-backend/internal/category/entities"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post/dtos"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/utils"
)

type PostRepo struct {
	Db *db.PrismaClient
}

func NewPostRepository(db *db.PrismaClient) post.PostRepository {
	return &PostRepo{
		Db: db,
	}
}

func (postRepo *PostRepo) CreatePost(cateResDat *entityCate.PostCategoryResDat, tagResDat *entityCate.PostTagResDat, req *dtos.PostReqDat) (*db.PostModel, error) {
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
		db.Post.Tags.Link(
			db.PostTag.ID.Equals(tagResDat.ID),
		),
		db.Post.Category.Link(
			db.PostCategory.ID.Equals(cateResDat.ID),
		),
	).With(
		db.Post.Category.Fetch(),
		db.Post.Tags.Fetch(),
		db.Post.Comments.Fetch(),
		db.Post.Likes.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}

	}
	return post, nil
}

func (postRepo *PostRepo) UpdatePostByUUID(cateResDat *entityCate.PostCategoryResDat, uuid string, req *dtos.PostReqDat) (*db.PostModel, error) {
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
		db.Post.Comments.Fetch(),
		db.Post.Likes.Fetch(),
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

func (postRepo *PostRepo) UpdatePostToPublisher(userUuid string, postUuid string) (string, error) {
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
		return "", &_errEntity.CError{
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
		return "", &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return pp.UUID, nil
}

func (postRepo *PostRepo) FetchPostByUUID(uuid string) (*db.PostModel, error) {
	ctx := context.Background()

	post, err := postRepo.Db.Post.FindUnique(
		db.Post.UUID.Equals(uuid),
	).With(
		db.Post.Category.Fetch(),
		db.Post.Tags.Fetch(),
		db.Post.Comments.Fetch(),
		db.Post.Likes.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return post, nil
}

func (postRepo *PostRepo) FetchPublisherPosts(opts *dtos.FetchPostOptReq) ([]db.PublicationPostModel, error) {
	ctx := context.Background()

	if opts.Page < 0 {
		opts.Page = 0
	}

	render := 10

	posts, err := postRepo.Db.PublicationPost.FindMany().With(
		db.PublicationPost.Post.Fetch().With(
			db.Post.Category.Fetch(),
			db.Post.Tags.Fetch(),
			db.Post.Comments.Fetch(),
			db.Post.Likes.Fetch(),
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
		db.Post.Comments.Fetch(),
		db.Post.Likes.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, &_errEntity.CError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return posts, nil
}

func (postRepo *PostRepo) DeletePostByUUID(postUuid string) error {
	ctx := context.Background()

	// unlink publisherPost and Delete if it present
	pDel, err := postRepo.Db.Post.FindUnique(
		db.Post.UUID.Equals(postUuid),
	).With(
		db.Post.Tags.Fetch(),
	).Delete().Exec(ctx)
	if err != nil {
		return &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if _, err := postRepo.Db.PostTag.FindUnique(
		db.PostTag.ID.Equals(pDel.Tags().ID),
	).Delete().Exec(ctx); err != nil {
		return &_errEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return nil
}
