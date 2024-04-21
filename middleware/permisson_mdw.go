package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionMdw() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuidParam := c.Param("user_uuid")
		if uuidParam == "" {
			c.JSON(400, gin.H{
				"status":      http.StatusText(http.StatusBadRequest),
				"status_code": http.StatusBadRequest,
				"message":     "User UUID is required",
				"result":      nil,
			})
			c.Abort()
			return
		}
		uuidHeader := c.GetHeader("user_uuid")

		if uuidParam != uuidHeader {
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
