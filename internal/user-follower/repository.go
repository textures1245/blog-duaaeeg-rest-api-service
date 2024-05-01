package userfollower

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user-follower/dtos"
)

type UserFollowerRepository interface {
	CreateUserFollower(usrFollowerUuid string, req *dtos.UserFollowerReqDat) (*db.UserFollowerModel, error)
	DeleteUserFollowerByUUID(usrFollowerUuid string, req *dtos.UserFollowerReqDat) error
}
