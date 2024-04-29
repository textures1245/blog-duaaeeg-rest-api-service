package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (routeRepo *RouteRepo) UserFollowerRoutes(spRoutes *gin.RouterGroup) {
	usrRg := spRoutes.Group("/users")

	usrFollowerRes := _userFollowerRepository.NewUsrFollowerRepo(routeRepo.Db)
	usrFollowerService := _userFollowerService.NewUsrFollowerService(usrFollowerRes)
	usrFollowerC := _userFollowerController.NewUsrFollowerController(usrFollowerService)

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
