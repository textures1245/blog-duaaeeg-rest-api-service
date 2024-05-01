package dtos

type LikeReqDat struct {
	UserUuid string `json:"user_uuid" db:"user_uuid" form:"user_uuid" binding:"required" validate:"required"`
}
