package dtos

type CommentReqDat struct {
	UserUuid string `json:"user_uuid" db:"user_uuid" form:"user_uuid" binding:"required" validate:"required"`
	Comment  string `json:"comment" db:"comment" form:"comment" binding:"required" validate:"required"`
}
