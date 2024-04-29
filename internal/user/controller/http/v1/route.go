package v1

import "github.com/gin-gonic/gin"

func (routeRepo *RouteRepo) UserRoutes(spRoutes *gin.RouterGroup) {
	usrRg := spRoutes.Group("/users")
	userRes := _userRepository.NewUserRepository(routeRepo.Db)
	userService := _userService.NewUserService(userRes)
	uC := _userController.NewUserController(userService)
	{

		usrRg.GET("/:user_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), uC.FetchUserByUUID)
		usrRg.POST("/:user_uuid/profile", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), uC.UpdateUserProfile)
	}
}
