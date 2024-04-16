package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/middleware"
	_authController "github.com/textures1245/BlogDuaaeeg-backend/model/auth/controller"
	_authRepository "github.com/textures1245/BlogDuaaeeg-backend/model/auth/repository"
	_authService "github.com/textures1245/BlogDuaaeeg-backend/model/auth/service"

	_userController "github.com/textures1245/BlogDuaaeeg-backend/model/user/controller"
	_userRepository "github.com/textures1245/BlogDuaaeeg-backend/model/user/repository"
	_userService "github.com/textures1245/BlogDuaaeeg-backend/model/user/service"

	_routeEntity "github.com/textures1245/BlogDuaaeeg-backend/routes/entity"
)

type RouteRepo struct {
	Db *db.PrismaClient
}

func (routeRepo *RouteRepo) RootRoutes(spRoutes *gin.RouterGroup) {
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

func NewRouteRepository(db *db.PrismaClient) _routeEntity.RouteRepository {
	return &RouteRepo{
		Db: db,
	}
}

// func (routeRepo *RouteRepo) PostsRoutes(spRoutes *gin.RouterGroup) {
// 	pRg := spRoutes.Group("/")
// 	{
// 		// userRouter.GET("/", spRoutes.UserControllers.GetUsers)

// 		// userRouter.POST("/", spRoutes.UserControllers.CreateUser)
// 	}

// }

// func (routeRepo *RouteRepo) AnalyticRoute(spRoutes *gin.RouterGroup) {
// 	analyticRg := spRoutes.Group("/analytics")
// 	{
// 		// userRouter.GET("/", spRoutes.UserControllers.GetUsers)

// 		// userRouter.POST("/", spRoutes.UserControllers.CreateUser)
// 	}
// }

// func (routeRepo *RouteRepo) PublicationsRoutes(spRoutes *gin.RouterGroup) {
// 	pubRg := spRoutes.Group("/")
// 	{
// 		// userRouter.GET("/", spRoutes.UserControllers.GetUsers)

// 		// userRouter.POST("/", spRoutes.UserControllers.CreateUser)
// 	}
// }

func (routeRepo *RouteRepo) UserRoutes(spRoutes *gin.RouterGroup) {
	usrRg := spRoutes.Group("/users")
	userRes := _userRepository.NewUserRepository(routeRepo.Db)
	userService := _userService.NewUserService(userRes)
	uC := _userController.NewUserController(userService)
	{
		usrRg.POST("/:user_uuid/profile", middleware.JwtAuthentication(), uC.UpdateUserProfile)
	}
}
