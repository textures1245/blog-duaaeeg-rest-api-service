package repository

import (
	"context"
	"net/http"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	errorEntity "github.com/textures1245/BlogDuaaeeg-backend/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/model/category/entity"
)

type cateRepo struct {
	Db *db.PrismaClient
}

func NewCateRepository(db *db.PrismaClient) entity.PostCategoryRepository {
	return &cateRepo{
		Db: db,
	}
}

func (c *cateRepo) CreateCategory(postUuid string, req *entity.PostCategoryReqDat) (*db.PostCategoryModel, error) {
	ctx := context.Background()

	cate, err := c.Db.PostCategory.CreateOne(
		db.PostCategory.Name.Set(req.Name),
		db.PostCategory.Post.Link(
			db.Post.UUID.Equals(postUuid),
		),
	).Exec(ctx)
	if err != nil {
		return nil, &errorEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return cate, nil
}

func (c *cateRepo) UpdateCategory(id int, req *entity.PostCategoryReqDat) (*db.PostCategoryModel, error) {
	ctx := context.Background()

	category, err := c.Db.PostCategory.FindUnique(
		db.PostCategory.ID.Equals(id),
	).Update(
		db.PostCategory.Name.Set(req.Name),
	).Exec(ctx)
	if err != nil {
		return nil, &errorEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return category, nil
}
