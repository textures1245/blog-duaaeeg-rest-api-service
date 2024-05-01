package category

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/category"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/category/dtos"
	errorEntity "github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
)

type cateRepo struct {
	Db *db.PrismaClient
}

func NewCateRepository(db *db.PrismaClient) category.PostCategoryRepository {
	return &cateRepo{
		Db: db,
	}
}

func (c *cateRepo) CreateOrUpdateCategory(req *dtos.PostCategoryReqDat) (*db.PostCategoryModel, error) {
	ctx := context.Background()

	capName := strings.Title(req.Name)

	oldCate, err := c.Db.PostCategory.FindUnique(
		db.PostCategory.Name.Equals(capName),
	).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, db.ErrNotFound) {
			newCate, err := c.Db.PostCategory.CreateOne(
				db.PostCategory.Name.Set(capName),
				// db.PostCategory.Post.Link(
				// 	db.Post.UUID.Equals(postUuid),
				// ),
			).Exec(ctx)
			if err != nil {
				return nil, &errorEntity.CError{
					StatusCode: http.StatusInternalServerError,
					Err:        err,
				}
			}
			return newCate, nil

		} else {
			return nil, &errorEntity.CError{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	// if _, err := c.Db.Post.FindUnique(
	// 	db.Post.UUID.Equals(postUuid),
	// ).Update(
	// 	db.Post.Category.Link(
	// 		db.PostCategory.ID.Equals(cate.ID),
	// 	),
	// ).Exec(ctx); err != nil {
	// 	return nil, &errorEntity.CError{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Err:        err,
	// 	}

	// }

	return oldCate, nil
}

// func (c *cateRepo) UpdateCategory(id int, req *entity.PostCategoryReqDat) (*db.PostCategoryModel, error) {
// 	ctx := context.Background()

// 	capName := strings.Title(req.Name)

// 	if _, err := c.Db.Post.FindUnique(
// 		db.Post.UUID.Equals(postUuid),
// 	).Update(
// 		db.Post.Category.Link(
// 			db.PostCategory.ID.Equals(id),
// 		),
// 	).Exec(ctx); err != nil {
// 		return nil, &errorEntity.CError{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}

// 	}

// 	cate, err := c.Db.PostCategory.FindUnique(
// 		db.PostCategory.ID.Equals(id),
// 	).Update(
// 		db.PostCategory.Name.Set(capName),
// 	).Exec(ctx)
// 	if err != nil {
// 		return nil, &errorEntity.CError{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}

// 	return cate, nil
// }
