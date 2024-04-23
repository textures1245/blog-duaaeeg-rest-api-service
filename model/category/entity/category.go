package entity

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type PostCategoryReqDat struct {
	Name string `json:"name" db:"name" form:"name" binding:"required" validate:"required"`
}

type PostCategoryResDat struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PostCategoryRepository interface {
	CreateOrUpdateCategory(postUuid string, req *PostCategoryReqDat) (*db.PostCategoryModel, error)
	// UpdateCategory(id int, req *PostCategoryReqDat) (*db.PostCategoryModel, error)
}
