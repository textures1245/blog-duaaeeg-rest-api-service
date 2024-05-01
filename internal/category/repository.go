package category

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/category/dtos"
)

type PostCategoryRepository interface {
	CreateOrUpdateCategory(req *dtos.PostCategoryReqDat) (*db.PostCategoryModel, error)
	// UpdateCategory(id int, req *PostCategoryReqDat) (*db.PostCategoryModel, error)
}

type PostTagRepository interface {
	CreateTags(req *dtos.PostTagReqDat) (*db.PostTagModel, error)
	UpdateTags(id int, req *dtos.PostTagReqDat) (*db.PostTagModel, error)
}
