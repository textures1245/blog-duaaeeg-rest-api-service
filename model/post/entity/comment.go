package entity

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type CommentReqDat struct {
	UserUuid string `json:"user_uuid" db:"user_uuid" form:"user_uuid" binding:"required" validate:"required"`
	Comment  string `json:"comment" db:"comment" form:"comment" binding:"required" validate:"required"`
}

type CommentResDat struct {
	UUID      string `json:"uuid"`
	UserUuid  string `json:"user_uuid"`
	Comment   string `json:"comment"`
	PostUUID  string `json:"post_uuid"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

type CommentRepository interface {
	CreateComment(postUuid string, req *CommentReqDat) (*db.CommentModel, error)
	UpdateCommentByUUID(cUuid string, req *CommentReqDat) (*db.CommentModel, error)
	DeleteCommentByUUID(cUuid string) error
}
