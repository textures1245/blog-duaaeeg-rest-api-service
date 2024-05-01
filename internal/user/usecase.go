package user

import (
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user/entities"
)

type UserService interface {
	OnFetchUserByUUID(userUuid string) (*entities.UserResDat, error)
	OnUpdateUserProfile(userUuid string, req *dtos.UserProfileDataRequest) (*entities.UserProfileRes, error)
}
