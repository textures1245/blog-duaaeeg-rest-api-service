package userfollower

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type UserFollowerRepository interface {
	CreateUserFollower(usrFollowerUuid string, req *UserFollowerReqDat) (*db.UserFollowerModel, error)
	DeleteUserFollowerByUUID(usrFollowerUuid string, req *UserFollowerReqDat) error
}
