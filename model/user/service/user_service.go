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

func (u *userUse) OnFetchUserByUUID(usrUuid string) (*entity.UserResDat, error) {
	user, err := u.userRepo.GetUserByUUID(usrUuid)

	if err != nil {
		return nil, err
	}

	res := &entity.UserResDat{
		UUID:  user.UUID,
		Email: user.Email,
		UserProfile: &entity.UserProfileRes{
			FirstName:      "",
			LastName:       "",
			Bio:            "",
			ProfilePicture: "",
			CreatedAt:      "",
			UpdateAt:       "",
		},
		Subscribers: user.UserFollowee(),
		Subscribing: user.UserFollowee(),
	}
	if usrProfile, ok := user.UserProfile(); ok {
		res.UserProfile = &entity.UserProfileRes{
			FirstName:      usrProfile.FirstName,
			LastName:       usrProfile.LastName,
			Bio:            usrProfile.Bio,
			ProfilePicture: usrProfile.ProfilePicture,
			CreatedAt:      usrProfile.CreatedAt.String(),
			UpdateAt:       usrProfile.UpdatedAt.String(),
		}
	}

	return res, nil

}

func (u *userUse) OnUpdateUserProfile(userUuid string, req *entity.UserProfileDataRequest) (*entity.UserProfileRes, error) {
	user, err := u.userRepo.UpdateProfile(userUuid, req)

	if err != nil {
		return nil, err
	}

	res := &entity.UserProfileRes{
		UUID:           user.UUID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Bio:            user.Bio,
		ProfilePicture: user.ProfilePicture,
		CreatedAt:      user.CreatedAt.String(),
		UpdateAt:       user.UpdatedAt.String(),
	}
	return res, nil
}
