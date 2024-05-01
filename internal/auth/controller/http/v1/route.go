package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/repository"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/auth/usecase"
	_userRepo "github.com/textures1245/BlogDuaaeeg-backend/internal/user/repository"
	_routeRepo "github.com/textures1245/BlogDuaaeeg-backend/pkg/datasource/route/repository"
	middleware "github.com/textures1245/BlogDuaaeeg-backend/pkg/middlewares"
)

type RouteRepo struct {
	*_routeRepo.RouteRepo
}

func (routeRepo *RouteRepo) AuthRoutes(spRoutes *gin.RouterGroup) {
	rootRg := spRoutes.Group("/")

	authRes := repository.NewAuthRepository(routeRepo.Db)
	userRes := _userRepo.NewUserRepository(routeRepo.Db)
	authService := usecase.NewAuthService(authRes, userRes)

	authC := NewAuthController(authService)

	{
		rootRg.POST("/login", authC.Login)

		rootRg.POST("/register", authC.Register)

		rootRg.GET("/auth-test", middleware.JwtAuthentication(), authC.AuthTest)
	}
}
