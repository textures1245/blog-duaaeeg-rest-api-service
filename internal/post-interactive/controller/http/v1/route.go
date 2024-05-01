package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive/repository/comment"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive/repository/like"
	_postInter "github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive/usecase"
	_routeRepo "github.com/textures1245/BlogDuaaeeg-backend/pkg/datasource/route/repository"
	middleware "github.com/textures1245/BlogDuaaeeg-backend/pkg/middlewares"
)

type RouteRepo struct {
	*_routeRepo.RouteRepo
}

func (routeRepo *RouteRepo) PostInteractiveRoutes(spRoutes *gin.RouterGroup) {
	pRg := spRoutes.Group("/post")

	commRepo := comment.NewCommRepo(routeRepo.Db)
	likeRepo := like.NewLikeRepo(routeRepo.Db)

	postInterService := _postInter.NewUserInteractiveUse(commRepo, likeRepo)
	postInterCtrl := NewPostInteractiveController(postInterService)

	// postInterCtrl := _postController.NewPostController(postService, cateService, usrInterService)

	// TODO: Test PostRoutes (DONE)
	{

		pRg.DELETE("/publish_posts/:post_uuid/:comment_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), postInterCtrl.UserDeleteComment)
		pRg.PATCH("/publish_posts/:post_uuid/:comment_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw(), postInterCtrl.UserUpdateComment)

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
				postInterCtrl.UserCommentedToPost(c)
			case "USER_LIKED":
				postInterCtrl.UserLikedPost(c)
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
				postInterCtrl.UserUnlikedPost(g)
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
