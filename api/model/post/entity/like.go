package entity

import "github.com/textures1245/BlogDuaaeeg-backend/api/db"

type LikeReqDat struct {
	UserUuid string `json:"user_uuid" db:"user_uuid" form:"user_uuid" binding:"required" validate:"required"`
}

type LikeResDat struct {
	UUID      string `json:"uuid"`
	UserUuid  string `json:"user_uuid"`
	PostUuid  string `json:"post_uuid"`
	CreatedAt string `json:"created_at"`
}

type LikeRepository interface {
	CreateLike(pUuid string, req *LikeReqDat) (*db.LikeModel, error)
	DeleteLikeByUUID(pUuid string, usrUuid string) error
}
