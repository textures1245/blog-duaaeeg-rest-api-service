package userfollower

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user-follower/dtos"
)

type UserFollowerService interface {
	OnSubscribeUser(usrFollowerUuid string, req *dtos.UserFollowerReqDat) (*db.UserFollowerModel, error)
	OnUnsubscribeUser(usrFollowerUuid string, req *dtos.UserFollowerReqDat) error
}
