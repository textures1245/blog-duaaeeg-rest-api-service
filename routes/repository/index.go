package repository

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/middleware"
	_authController "github.com/textures1245/BlogDuaaeeg-backend/model/auth/controller"
	_authRepository "github.com/textures1245/BlogDuaaeeg-backend/model/auth/repository"
	_authService "github.com/textures1245/BlogDuaaeeg-backend/model/auth/service"

	_userController "github.com/textures1245/BlogDuaaeeg-backend/model/user/controller"
	_userRepository "github.com/textures1245/BlogDuaaeeg-backend/model/user/repository"
	_userService "github.com/textures1245/BlogDuaaeeg-backend/model/user/service"

	_postController "github.com/textures1245/BlogDuaaeeg-backend/model/post/controller"
	_postRepository "github.com/textures1245/BlogDuaaeeg-backend/model/post/repository"
	_postService "github.com/textures1245/BlogDuaaeeg-backend/model/post/service"

	_cateRepository "github.com/textures1245/BlogDuaaeeg-backend/model/category/repository"
	_cateService "github.com/textures1245/BlogDuaaeeg-backend/model/category/service"

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

func (routeRepo *RouteRepo) PostsRoutes(spRoutes *gin.RouterGroup) {
	pRg := spRoutes.Group("/post")
	postRes := _postRepository.NewPostRepository(routeRepo.Db)
	userRes := _userRepository.NewUserRepository(routeRepo.Db)
	tagRepo := _cateRepository.NewTagRepository(routeRepo.Db)
	postService := _postService.NewPostService(postRes, userRes, tagRepo)

	cateRes := _cateRepository.NewCateRepository(routeRepo.Db)
	tagRes := _cateRepository.NewTagRepository(routeRepo.Db)
	cateService := _cateService.NewCategoryService(cateRes, tagRes)
	pC := _postController.NewPostController(postService, cateService)

	{
		pRg.GET("/publish_posts", middleware.JwtAuthentication(), pC.GetPublisherPosts)
		pRg.GET("/publish_posts/:post_uuid", middleware.JwtAuthentication(), pC.GetPostByUUID)
		pRg.GET("/:user_uuid/posts/:post_uuid", middleware.JwtAuthentication(), pC.GetPostByUUID)
		pRg.GET("/:user_uuid/posts", middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.GetPostByUserUUID)
		pRg.POST("/:user_uuid/post_form", middleware.JwtAuthentication(), middleware.PermissionMdw(), func(c *gin.Context) {
			a := c.Query("action")
			if a == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "missing required query params",
					"result":      nil,
				})
				return
			}

			switch a {
			case "CREATE":
				pC.CreatePost(c)
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "action query params is invalid",
					"result":      nil,
				})
			}

		})
		pRg.PATCH("/:user_uuid/post_form/:post_uuid", middleware.JwtAuthentication(), middleware.PermissionMdw(), func(c *gin.Context) {
			a := c.Query("action")
			if a == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "missing required query params",
					"result":      nil,
				})
				return
			}

			switch a {
			case "UPDATE":
				pC.UpdatePost(c)
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "action query params is invalid",
					"result":      nil,
				})
			}
		})
		pRg.DELETE("/:user_uuid/post_form/:post_uuid", middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.DeletePostAndPublisherPostByUUID)

	}
}

// TODO: Test PostRoutes

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
		usrRg.POST("/:user_uuid/profile", middleware.JwtAuthentication(), middleware.PermissionMdw(), uC.UpdateUserProfile)
	}
}
