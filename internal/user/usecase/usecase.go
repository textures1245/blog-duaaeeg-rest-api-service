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

func (u *userUse) OnFetchUsers() ([]*entities.UserResDat, error) {
	users, err := u.userRepo.GetUsers()

	if err != nil {
		return nil, err
	}

	var res []*entities.UserResDat
	for _, user := range users {
		usr := &entities.UserResDat{
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
			CreatedAt:   user.CreatedAt.String(),
			UpdatedAt:   user.UpdatedAt.String(),
		}
		if usrProfile, ok := user.UserProfile(); ok {
			usr.UserProfile = &entities.UserProfileRes{
				FirstName:      usrProfile.FirstName,
				LastName:       usrProfile.LastName,
				Bio:            usrProfile.Bio,
				ProfilePicture: usrProfile.ProfilePicture,
				CreatedAt:      usrProfile.CreatedAt.String(),
				UpdateAt:       usrProfile.UpdatedAt.String(),
			}
		}
		res = append(res, usr)
	}

	return res, nil
}

func (u *userUse) OnFetchUsersWithPW() ([]*entities.UserWithPWResDat, error) {
	users, err := u.userRepo.GetUsers()

	if err != nil {
		return nil, err
	}

	var res []*entities.UserWithPWResDat
	for _, user := range users {
		usr := &entities.UserWithPWResDat{
			UUID:     user.UUID,
			Email:    user.Email,
			Password: user.Password,
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
			CreatedAt:   user.CreatedAt.String(),
			UpdatedAt:   user.UpdatedAt.String(),
		}
		if usrProfile, ok := user.UserProfile(); ok {
			usr.UserProfile = &entities.UserProfileRes{
				FirstName:      usrProfile.FirstName,
				LastName:       usrProfile.LastName,
				Bio:            usrProfile.Bio,
				ProfilePicture: usrProfile.ProfilePicture,
				CreatedAt:      usrProfile.CreatedAt.String(),
				UpdateAt:       usrProfile.UpdatedAt.String(),
			}
		}
		res = append(res, usr)
	}

	return res, nil
}

func (u *userUse) OnDeleteUser(userUuid string) error {
	err := u.userRepo.DeleteUserByUuid(userUuid)

	if err != nil {
		return err
	}

	return nil
}
