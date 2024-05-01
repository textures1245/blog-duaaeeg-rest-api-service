package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user-follower/repository"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/user-follower/usecase"
	_routeRepo "github.com/textures1245/BlogDuaaeeg-backend/pkg/datasource/route/repository"
	middleware "github.com/textures1245/BlogDuaaeeg-backend/pkg/middlewares"
)

type RouteRepo struct {
	*_routeRepo.RouteRepo
}

func (routeRepo *RouteRepo) UserFollowerRoutes(spRoutes *gin.RouterGroup) {
	usrRg := spRoutes.Group("/users")

	usrFollowerRes := repository.NewUsrFollowerRepo(routeRepo.Db)
	usrFollowerService := usecase.NewUsrFollowerService(usrFollowerRes)
	usrFollowerC := NewUsrFollowerController(usrFollowerService)

	{
		usrRg.POST("/:user_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), middleware.PermissionMdw([]string{"OWNER_ACTION_FORBIDDEN", "PREVENT_DEFAULT_ACTION"}), func(c *gin.Context) {
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
		usrRg.DELETE("/:user_uuid", middleware.CORSConfig(), middleware.JwtAuthentication(), func(c *gin.Context) {
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
