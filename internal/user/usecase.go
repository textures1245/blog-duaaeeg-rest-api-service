package user

import (
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user/entities"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/utils"
)

type UserService interface {
	OnFetchUserByUUID(userUuid string) (*entities.UserResDat, error)
	OnUpdateUserProfile(userUuid string, req *dtos.UserProfileDataRequest) (*entities.UserProfileRes, error)
	OnFetchUsers() ([]*entities.UserResDat, error)
	OnFetchUsersWithPW() ([]*entities.UserWithPWResDat, error)
	OnDeleteUser(userUuid string) error
	OnExportToExcel() (*utils.ExcelData, error)
}
