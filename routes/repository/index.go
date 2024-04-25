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

	_userFollowerController "github.com/textures1245/BlogDuaaeeg-backend/model/user-follower/controller"
	_userFollowerRepository "github.com/textures1245/BlogDuaaeeg-backend/model/user-follower/repository"
	_userFollowerService "github.com/textures1245/BlogDuaaeeg-backend/model/user-follower/service"

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
	commRepo := _postRepository.NewCommRepo(routeRepo.Db)
	likeRepo := _postRepository.NewLikeRepo(routeRepo.Db)

	postService := _postService.NewPostService(postRes, userRes, tagRepo)

	usrInterService := _postService.NewUserInteractiveUse(commRepo, likeRepo)
	cateRes := _cateRepository.NewCateRepository(routeRepo.Db)
	tagRes := _cateRepository.NewTagRepository(routeRepo.Db)
	cateService := _cateService.NewCategoryService(cateRes, tagRes)
	pC := _postController.NewPostController(postService, cateService, usrInterService)

	// TODO: Test PostRoutes (DONE)
	{
		pRg.GET("/publish_posts", middleware.JwtAuthentication(), pC.GetPublisherPosts)
		pRg.GET("/publish_posts/:post_uuid", middleware.JwtAuthentication(), pC.GetPostByUUID)
		pRg.GET("/:user_uuid/posts/:post_uuid", middleware.JwtAuthentication(), pC.GetPostByUUID)
		pRg.GET("/:user_uuid/posts", middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.GetPostByUserUUID)
		pRg.POST("/:user_uuid/post_form", middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.CreatePost)

		// TODO: Test UserInteractiveRoutes
		pRg.PATCH("/:user_uuid/post_form/:post_uuid", middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.UpdatePost)

		pRg.DELETE("/:user_uuid/post_form/:post_uuid", middleware.JwtAuthentication(), middleware.PermissionMdw(), func(c *gin.Context) {
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
			case "OWNER_DELETE_POST":
				pC.DeletePostAndPublisherPostByUUID(c)
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "action query params is invalid",
					"result":      nil,
				})
			}

		})

		pRg.DELETE("/publish_posts/:post_uuid/:comment_uuid", middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.UserDeleteComment)
		pRg.PATCH("/publish_posts/:post_uuid/:comment_uuid", middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.UserUpdateComment)

		pRg.POST("/publish_posts/:post_uuid", middleware.JwtAuthentication(), func(c *gin.Context) {
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
			case "USER_COMMENTED":
				pC.UserCommentedToPost(c)
			case "USER_LIKED":
				pC.UserLikedPost(c)
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "action query params is invalid",
					"result":      nil,
				})
			}
		})
		pRg.DELETE("/publish_posts/:post_uuid", middleware.JwtAuthentication(), func(g *gin.Context) {
			a := g.Query("action")
			if a == "" {
				g.JSON(http.StatusBadRequest, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "missing required query params",
					"result":      nil,
				})
				return
			}

			switch a {
			case "USER_UNLIKED":
				pC.UserUnlikedPost(g)
			default:
				g.JSON(http.StatusBadRequest, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "action query params is invalid",
					"result":      nil,
				})
			}

		})
	}
}

func (routeRepo *RouteRepo) UserRoutes(spRoutes *gin.RouterGroup) {
	usrRg := spRoutes.Group("/users")
	userRes := _userRepository.NewUserRepository(routeRepo.Db)
	userService := _userService.NewUserService(userRes)
	uC := _userController.NewUserController(userService)

	usrFollowerRes := _userFollowerRepository.NewUsrFollowerRepo(routeRepo.Db)
	usrFollowerService := _userFollowerService.NewUsrFollowerService(usrFollowerRes)
	usrFollowerC := _userFollowerController.NewUsrFollowerController(usrFollowerService)

	{
		usrRg.POST("/:user_uuid/profile", middleware.JwtAuthentication(), middleware.PermissionMdw(), uC.UpdateUserProfile)

		usrRg.POST("/:user_uuid", middleware.JwtAuthentication(), middleware.PermissionMdw([]string{"OWNER_ACTION_FORBIDDEN", "PREVENT_DEFAULT_ACTION"}), func(c *gin.Context) {
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
			case "USER_SUBSCRIBE":
				usrFollowerC.SubscribeUser(c)
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "action query params is invalid",
					"result":      nil,
				})
			}

		})
		usrRg.DELETE("/:user_uuid", middleware.JwtAuthentication(), func(c *gin.Context) {
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
			case "USER_UNSUBSCRIBE":
				usrFollowerC.UnsubscribeUser(c)
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "action query params is invalid",
					"result":      nil,
				})
			}

		})
	}
}

// TODO: Implemented AnalyticRoute
// func (routeRepo *RouteRepo) AnalyticRoute(spRoutes *gin.RouterGroup) {
// 	analyticRg := spRoutes.Group("/analytics")
// 	{
// 		// userRouter.GET("/", spRoutes.UserControllers.GetUsers)

// 		// userRouter.POST("/", spRoutes.UserControllers.CreateUser)
// 	}
// }
