package entity

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type UserFollowerReqDat struct {
	UserFolloweeUuid string `json:"user_followee_uuid" db:"user_followee_uuid" form:"user_followee_uuid" binding:"required" validate:"required"`
}

// type UserFollowerResDat struct {
// 	ID               string `json:"id"`
// 	UserFollowerUuid string `json:"user_follower_uuid"`
// 	UserFolloweeUuid string `json:"user_followee_uuid"`
// 	CreatedAt        string `json:"created_at"`
// 	UpdateAt         string `json:"update_at"`
// }

type UserFollowerRepository interface {
	CreateUserFollower(usrFollowerUuid string, req *UserFollowerReqDat) (*db.UserFollowerModel, error)
	DeleteUserFollowerByUUID(usrFollowerUuid string, req *UserFollowerReqDat) error
}

type UserFollowerService interface {
	OnSubscribeUser(usrFollowerUuid string, req *UserFollowerReqDat) (*db.UserFollowerModel, error)
	OnUnsubscribeUser(usrFollowerUuid string, req *UserFollowerReqDat) error
}
