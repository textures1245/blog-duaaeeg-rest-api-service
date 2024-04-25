package service

import (
	"github.com/textures1245/BlogDuaaeeg-backend/api/db"
	"github.com/textures1245/BlogDuaaeeg-backend/api/model/user-follower/entity"
)

type usrFollowerUse struct {
	usrFollowerRepo entity.UserFollowerRepository
}

func NewUsrFollowerService(usrFollowerRepo entity.UserFollowerRepository) entity.UserFollowerService {
	return &usrFollowerUse{
		usrFollowerRepo,
	}
}

func (u *usrFollowerUse) OnSubscribeUser(usrFollowerUuid string, req *entity.UserFollowerReqDat) (*db.UserFollowerModel, error) {
	usrFollower, err := u.usrFollowerRepo.CreateUserFollower(usrFollowerUuid, req)
	if err != nil {
		return nil, err
	}
	return usrFollower, nil
}

func (u *usrFollowerUse) OnUnsubscribeUser(usrFollowerUuid string, req *entity.UserFollowerReqDat) error {
	err := u.usrFollowerRepo.DeleteUserFollowerByUUID(usrFollowerUuid, req)
	if err != nil {
		return err
	}
	return nil
}
