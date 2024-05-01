package usecase

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	_usrFollower "github.com/textures1245/BlogDuaaeeg-backend/internal/user-follower"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user-follower/dtos"
)

type usrFollowerUse struct {
	usrFollowerRepo _usrFollower.UserFollowerRepository
}

func NewUsrFollowerService(usrFollowerRepo _usrFollower.UserFollowerRepository) _usrFollower.UserFollowerService {
	return &usrFollowerUse{
		usrFollowerRepo,
	}
}

func (u *usrFollowerUse) OnSubscribeUser(usrFollowerUuid string, req *dtos.UserFollowerReqDat) (*db.UserFollowerModel, error) {
	usrFollower, err := u.usrFollowerRepo.CreateUserFollower(usrFollowerUuid, req)
	if err != nil {
		return nil, err
	}
	return usrFollower, nil
}

func (u *usrFollowerUse) OnUnsubscribeUser(usrFollowerUuid string, req *dtos.UserFollowerReqDat) error {
	err := u.usrFollowerRepo.DeleteUserFollowerByUUID(usrFollowerUuid, req)
	if err != nil {
		return err
	}
	return nil
}
