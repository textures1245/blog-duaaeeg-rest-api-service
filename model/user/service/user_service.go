package service

import (
	"github.com/textures1245/BlogDuaaeeg-backend/model/user/entity"
)

type userUse struct {
	userRepo entity.UsersRepository
}

func NewUserService(userRepo entity.UsersRepository) entity.UserService {
	return &userUse{
		userRepo: userRepo,
	}
}

func (u *userUse) OnUpdateUserProfile(req *entity.UserProfileDataRequest) (*entity.UserProfileRes, error) {
	user, err := u.userRepo.UpdateProfile(req)
	if err != nil {
		return nil, err
	}

	res := &entity.UserProfileRes{
		UUID:           user.UUID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Bio:            user.Bio,
		ProfilePicture: user.ProfilePicture,
	}
	return res, nil
}
