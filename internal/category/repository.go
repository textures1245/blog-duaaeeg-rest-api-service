package category

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type PostCategoryRepository interface {
	CreateOrUpdateCategory(req *PostCategoryReqDat) (*db.PostCategoryModel, error)
	// UpdateCategory(id int, req *PostCategoryReqDat) (*db.PostCategoryModel, error)
}

type PostTagRepository interface {
	CreateTags(req *PostTagReqDat) (*db.PostTagModel, error)
	UpdateTags(id int, req *PostTagReqDat) (*db.PostTagModel, error)
}
