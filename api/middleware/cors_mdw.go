package middleware

import (
	"github.com/gin-gonic/gin"
)

type CORSOption struct {
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	AllowCredentials string
}

func CORSConfig(opt ...*CORSOption) gin.HandlerFunc {

	defaultAllowOrigins := "*"
	defaultAllowMethods := "GET, POST, PUT, DELETE, OPTIONS"
	defaultAllowHeaders := "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"
	defaultAllowCredentials := "true"

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", defaultAllowOrigins)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", defaultAllowCredentials)
		c.Writer.Header().Set("Access-Control-Allow-Methods", defaultAllowMethods)
		c.Writer.Header().Set("Access-Control-Allow-Headers", defaultAllowHeaders)

		if len(opt) > 0 {
			if opt[0].AllowOrigins != "" {
				c.Writer.Header().Set("Access-Control-Allow-Origin", opt[0].AllowOrigins)
			}
			if opt[0].AllowMethods != "" {
				c.Writer.Header().Set("Access-Control-Allow-Methods", opt[0].AllowMethods)
			}
			if opt[0].AllowHeaders != "" {
				c.Writer.Header().Set("Access-Control-Allow-Headers", opt[0].AllowHeaders)
			}
			if opt[0].AllowCredentials != "" {
				c.Writer.Header().Set("Access-Control-Allow-Credentials", opt[0].AllowCredentials)
			}
		}
		c.Next()
	}
}

// TODO: Implement more secure on CORS middleware in route handler
