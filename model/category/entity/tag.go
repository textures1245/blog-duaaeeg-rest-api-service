package entity

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type PostTagReqDat struct {
	Tags []string `json:"tags" db:"tag" form:"tag" binding:"required" validate:"required"`
}

type PostTagResDat struct {
	ID   int      `json:"id"`
	Tags []string `json:"tags"`
}

type PostTagRepository interface {
	CreateTags(req *PostTagReqDat) (*db.PostTagModel, error)
	UpdateTags(id int, req *PostTagReqDat) (*db.PostTagModel, error)
}
