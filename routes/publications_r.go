package routes

import (
	"github.com/gin-gonic/gin"
)

func publicationsRoutes(spRoutes *gin.RouterGroup) {
	pubRg := spRoutes.Group("/")
	{
		// userRouter.GET("/", spRoutes.UserControllers.GetUsers)

		// userRouter.POST("/", spRoutes.UserControllers.CreateUser)
	}
}
