package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(ginRoute *gin.RouterGroup) {
	userRoutes(ginRoute)
	postsRoutes(ginRoute)
	publicationsRoutes(ginRoute)
	rootRoutes(ginRoute)
}
