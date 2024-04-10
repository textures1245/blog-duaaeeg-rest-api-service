package routes

import (
	"github.com/gin-gonic/gin"
)

func userRoutes(spRoutes *gin.RouterGroup) {
	usrRg := spRoutes.Group("/users")
	{
		// userRouter.GET("/", spRoutes.UserControllers.GetUsers)

		// userRouter.POST("/", spRoutes.UserControllers.CreateUser)
	}
}
