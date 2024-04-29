package route

import "github.com/gin-gonic/gin"

type RouteRepository interface {
	AuthRoutes(spRoutes *gin.RouterGroup)
	PostsRoutes(spRoutes *gin.RouterGroup)
	// PublicationsRoutes(spRoutes *gin.RouterGroup)
	// AnalyticRoute(spRoutes *gin.RouterGroup)
	UserRoutes(spRoutes *gin.RouterGroup)
}
