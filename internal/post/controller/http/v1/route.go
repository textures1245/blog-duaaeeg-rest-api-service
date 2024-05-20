package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_cateRepo "github.com/textures1245/BlogDuaaeeg-backend/internal/category/repository/category"
	_tagRepo "github.com/textures1245/BlogDuaaeeg-backend/internal/category/repository/tag"
	_cateUsecase "github.com/textures1245/BlogDuaaeeg-backend/internal/category/usecase"
	_fileRepo "github.com/textures1245/BlogDuaaeeg-backend/internal/file/repository"
	_fileUsecase "github.com/textures1245/BlogDuaaeeg-backend/internal/file/usecase"
	_postRepo "github.com/textures1245/BlogDuaaeeg-backend/internal/post/repository"
	_postUsecase "github.com/textures1245/BlogDuaaeeg-backend/internal/post/usecase"
	_userRepo "github.com/textures1245/BlogDuaaeeg-backend/internal/user/repository"
	_routeRepo "github.com/textures1245/BlogDuaaeeg-backend/pkg/datasource/route/repository"
	middleware "github.com/textures1245/BlogDuaaeeg-backend/pkg/middlewares"
)

type RouteRepo struct {
	*_routeRepo.RouteRepo
}

func (routeRepo *RouteRepo) PostsRoutes(spRoutes *gin.RouterGroup) {
	pRg := spRoutes.Group("/post")
	postRes := _postRepo.NewPostRepository(routeRepo.Db)
	userRes := _userRepo.NewUserRepository(routeRepo.Db)
	tagRepo := _tagRepo.NewTagRepository(routeRepo.Db)
	cateRepo := _cateRepo.NewCateRepository(routeRepo.Db)

	fileRepo := _fileRepo.NewFileRepository(routeRepo.Db)
	fileUsecase := _fileUsecase.NewFileUsecase(fileRepo)

	postUsecase := _postUsecase.NewPostService(postRes, userRes, tagRepo, fileUsecase)
	cateUsecase := _cateUsecase.NewCategoryService(cateRepo, tagRepo)

	postCtrl := NewPostController(postUsecase, cateUsecase)

	// TODO: Test PostRoutes (DONE)
	{
		pRg.GET("/publish_posts", middleware.CORSConfig(), middleware.JwtAuthentication(), postCtrl.GetPublisherPosts)
		pRg.GET("/publish_posts/:post_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), postCtrl.GetPostByUUID)
		pRg.GET("/:user_uuid/posts/:post_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), postCtrl.GetPostByUUID)
		pRg.GET("/:user_uuid/posts", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), postCtrl.GetPostByUserUUID)
		pRg.POST("/:user_uuid/post_form", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), postCtrl.CreatePost)

		// TODO: Test UserInteractiveRoutes
		pRg.PATCH("/:user_uuid/post_form/:post_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), postCtrl.UpdatePost)

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
				postCtrl.DeletePostAndPublisherPostByUUID(c)
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
