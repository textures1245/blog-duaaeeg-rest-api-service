package v1

import (
	"github.com/gin-gonic/gin"
	_userRepository "github.com/textures1245/BlogDuaaeeg-backend/internal/user/repository"
	_userUsecase "github.com/textures1245/BlogDuaaeeg-backend/internal/user/usecase"
	_routeRepo "github.com/textures1245/BlogDuaaeeg-backend/pkg/datasource/route/repository"
	middleware "github.com/textures1245/BlogDuaaeeg-backend/pkg/middlewares"
)

type RouteRepo struct {
	*_routeRepo.RouteRepo
}

func (routeRepo *RouteRepo) UserRoutes(spRoutes *gin.RouterGroup) {
	usrRg := spRoutes.Group("/users")
	userRes := _userRepository.NewUserRepository(routeRepo.Db)
	userService := _userUsecase.NewUserService(userRes)
	uC := NewUserController(userService)
	{

		usrRg.GET("/:user_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), uC.FetchUserByUUID)
		usrRg.POST("/:user_uuid/profile", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), uC.UpdateUserProfile)
	}
}