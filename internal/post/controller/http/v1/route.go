package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_cateRepo "github.com/textures1245/BlogDuaaeeg-backend/internal/category/repository/category"
	_tagRepo "github.com/textures1245/BlogDuaaeeg-backend/internal/category/repository/tag"
	_cateUsecase "github.com/textures1245/BlogDuaaeeg-backend/internal/category/usecase"

	v1 "github.com/textures1245/BlogDuaaeeg-backend/internal/post/controller/http/v1"
)

func (routeRepo *RouteRepo) PostsRoutes(spRoutes *gin.RouterGroup) {
	pRg := spRoutes.Group("/post")
	postRes := _postRepository.NewPostRepository(routeRepo.Db)
	userRes := _userRepository.NewUserRepository(routeRepo.Db)
	tagRepo := _tagRepo.NewTagRepository(routeRepo.Db)
	cateRepo := _cateRepo.NewCateRepository(routeRepo.Db)

	postService := _postService.NewPostService(postRes, userRes, tagRepo)
	cateService := _cateUsecase.NewCategoryService(cateRepo, tagRepo)

	postCtrl := v1.NewPostController(postService, postInterService)

	// TODO: Test PostRoutes (DONE)
	{
		pRg.GET("/publish_posts", middleware.CORSConfig(), middleware.JwtAuthentication(), pC.GetPublisherPosts)
		pRg.GET("/publish_posts/:post_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), pC.GetPostByUUID)
		pRg.GET("/:user_uuid/posts/:post_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), pC.GetPostByUUID)
		pRg.GET("/:user_uuid/posts", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.GetPostByUserUUID)
		pRg.POST("/:user_uuid/post_form", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.CreatePost)

		// TODO: Test UserInteractiveRoutes
		pRg.PATCH("/:user_uuid/post_form/:post_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.UpdatePost)

		pRg.DELETE("/:user_uuid/post_form/:post_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), func(c *gin.Context) {
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

		pRg.DELETE("/publish_posts/:post_uuid/:comment_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.UserDeleteComment)
		pRg.PATCH("/publish_posts/:post_uuid/:comment_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), pC.UserUpdateComment)

		pRg.POST("/publish_posts/:post_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), func(c *gin.Context) {
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
		pRg.DELETE("/publish_posts/:post_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), func(g *gin.Context) {
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
