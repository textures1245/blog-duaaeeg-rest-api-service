package usecase

import (
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user/entities"
)

type userUse struct {
	userRepo user.UsersRepository
}

func NewUserService(userRepo user.UsersRepository) user.UserService {
	return &userUse{
		userRepo: userRepo,
	}
}

func (u *userUse) OnFetchUserByUUID(usrUuid string) (*entities.UserResDat, error) {
	user, err := u.userRepo.GetUserByUUID(usrUuid)

	if err != nil {
		return nil, err
	}

	res := &entities.UserResDat{
		UUID:  user.UUID,
		Email: user.Email,
		UserProfile: &entities.UserProfileRes{
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
		res.UserProfile = &entities.UserProfileRes{
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

func (u *userUse) OnUpdateUserProfile(userUuid string, req *dtos.UserProfileDataRequest) (*entities.UserProfileRes, error) {
	user, err := u.userRepo.UpdateProfile(userUuid, req)

	if err != nil {
		return nil, err
	}

	res := &entities.UserProfileRes{
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
