package entity

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type LikeReqDat struct {
	UserUuid string `json:"user_uuid" db:"user_uuid" form:"user_uuid" binding:"required" validate:"required"`
	PostUuid string `json:"post_uuid" db:"post_uuid" form:"post_uuid" binding:"required" validate:"required"`
}

type LikeResDat struct {
	UUID      string `json:"uuid"`
	UserUuid  string `json:"user_uuid"`
	PostUuid  string `json:"post_uuid"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

type LikeRepository interface {
	CreateLike(req *LikeReqDat) (*db.LikeModel, error)
	DeleteLikeByUUID(uuid string) error
}

// type LikeService interface {
// 	OnLikedPost(req *LikeReqDat) (*LikeResDat, error)
// 	OnUnlikedPost(uuid string) error
// }
