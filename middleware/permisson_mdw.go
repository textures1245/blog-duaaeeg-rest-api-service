package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/model/utils"
)

func PermissionMdw(opt ...[]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuidParam := c.Param("user_uuid")
		if uuidParam == "" {
			uuidBind := struct {
				UserUUID string `json:"user_uuid" form:"user_uuid" binding:"required"`
			}{}
			if err := c.ShouldBind(&uuidBind); err != nil {
				c.JSON(400, gin.H{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "User UUID is required",
					"result":      nil,
				})
				c.Abort()
				return
			}
			fmt.Println(uuidBind.UserUUID)
			uuidParam = uuidBind.UserUUID
		}
		uuidC := c.MustGet("user_uuid")

		if uuidParam != uuidC {
			if len(opt) > 0 {
				opt := opt[0]
				if utils.Contains(opt, "PREVENT_DEFAULT_ACTION") {
					c.Next()
					c.Abort()
					return
				}
			}
			c.JSON(403, gin.H{
				"status":      http.StatusText(http.StatusForbidden),
				"status_code": http.StatusForbidden,
				"message":     "You don't have permission to access this resource",
				"result":      nil,
			})
			c.Abort()
			return
		} else {
			if len(opt) > 0 {
				opt := opt[0]
				if utils.Contains(opt, "OWNER_ACTION_FORBIDDEN") {
					c.JSON(403, gin.H{
						"status":      http.StatusText(http.StatusForbidden),
						"status_code": http.StatusForbidden,
						"message":     "Owner action is forbidden",
						"result":      nil,
					})
					c.Abort()
					return
				}
			}
			c.Next()
		}
	}
}
