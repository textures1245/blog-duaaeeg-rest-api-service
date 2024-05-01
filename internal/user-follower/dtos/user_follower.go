package dtos

type UserFollowerReqDat struct {
	UserFolloweeUuid string `json:"user_followee_uuid" db:"user_followee_uuid" form:"user_followee_uuid" binding:"required" validate:"required"`
}
