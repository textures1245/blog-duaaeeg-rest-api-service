package userfollower

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type UserFollowerService interface {
	OnSubscribeUser(usrFollowerUuid string, req *UserFollowerReqDat) (*db.UserFollowerModel, error)
	OnUnsubscribeUser(usrFollowerUuid string, req *UserFollowerReqDat) error
}
