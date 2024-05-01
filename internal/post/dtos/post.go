package dtos

import (
	dtos "github.com/textures1245/BlogDuaaeeg-backend/internal/category/dtos"
)

type PostReqDat struct {
	UserUuid     string                   `json:"user_uuid" db:"user_uuid" form:"user_uuid" binding:"required" validate:"required"`
	Title        string                   `json:"title" db:"title" form:"title" binding:"required" validate:"required"`
	Content      string                   `json:"content" db:"content" form:"content" binding:"required" validate:"required"`
	Published    bool                     `json:"published" db:"published" form:"published"`
	SrcType      string                   `json:"src_type" db:"src_type" form:"src_type" binding:"required" validate:"required"`
	PostCategory *dtos.PostCategoryReqDat `json:"category" db:"category" form:"category" binding:"required"`
	PostTag      *dtos.PostTagReqDat      `json:"tags" db:"tags" form:"tags" binding:"required"`
}
