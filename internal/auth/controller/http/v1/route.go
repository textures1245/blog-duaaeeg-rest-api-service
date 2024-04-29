package v1

import (
	"github.com/gin-gonic/gin"
	_routeRepo "github.com/textures1245/BlogDuaaeeg-backend/pkg/datasource/route/repository"
)

func (routeRepo *_routeRepo.RouteRepo) AuthRoutes(spRoutes *gin.RouterGroup) {
	rootRg := spRoutes.Group("/")

	authRes := _authRepository.NewAuthRepository(routeRepo.Db)
	userRes := _userRepository.NewUserRepository(routeRepo.Db)
	authService := _authService.NewAuthService(authRes, userRes)

	authC := _authController.NewAuthController(authService)

	{
		rootRg.POST("/login", authC.Login)

		rootRg.POST("/register", authC.Register)

		rootRg.GET("/auth-test", middleware.JwtAuthentication(), authC.AuthTest)
	}
}
