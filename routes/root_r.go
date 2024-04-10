package routes

import (
	"github.com/gin-gonic/gin"
)

func rootRoutes(spRoutes *gin.RouterGroup) {
	rootRg := spRoutes.Group("/")
	{
		// userRouter.GET("/", spRoutes.UserControllers.GetUsers)

		// userRouter.POST("/", spRoutes.UserControllers.CreateUser)
	}
}
