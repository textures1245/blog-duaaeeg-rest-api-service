package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionMdw() gin.HandlerFunc {
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
			c.JSON(403, gin.H{
				"status":      http.StatusText(http.StatusForbidden),
				"status_code": http.StatusForbidden,
				"message":     "You don't have permission to access this resource",
				"result":      nil,
			})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
