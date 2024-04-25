package repository

import (
	"context"
	"net/http"

	"github.com/textures1245/BlogDuaaeeg-backend/api/db"
	errorEntity "github.com/textures1245/BlogDuaaeeg-backend/api/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/api/model/category/entity"
)

type tagRepo struct {
	Db *db.PrismaClient
}

func NewTagRepository(db *db.PrismaClient) entity.PostTagRepository {
	return &tagRepo{
		Db: db,
	}
}

func (t *tagRepo) UpdateTags(id int, req *entity.PostTagReqDat) (*db.PostTagModel, error) {
	ctx := context.Background()

	postTag, err := t.Db.PostTag.FindUnique(
		db.PostTag.ID.Equals(id),
	).Update(
		db.PostTag.Tags.Set(req.Tags),
	).Exec(ctx)
	if err != nil {
		return nil, &errorEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return postTag, nil
}

func (t *tagRepo) CreateTags(req *entity.PostTagReqDat) (*db.PostTagModel, error) {
	ctx := context.Background()

	p, err := t.Db.PostTag.CreateOne().Exec(ctx)

	if err != nil {
		return nil, &errorEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	postTag, err := t.UpdateTags(p.ID, req)

	if err != nil {
		return nil, &errorEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return postTag, nil
}
