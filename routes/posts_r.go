package routes

import (
	"github.com/gin-gonic/gin"
)

func postsRoutes(spRoutes *gin.RouterGroup) {
	pRg := spRoutes.Group("/")
	{
		// userRouter.GET("/", spRoutes.UserControllers.GetUsers)

		// userRouter.POST("/", spRoutes.UserControllers.CreateUser)
	}
}
