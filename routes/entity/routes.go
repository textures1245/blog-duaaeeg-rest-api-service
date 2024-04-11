package entity

import "github.com/gin-gonic/gin"

type RouteRepository interface {
	RootRoutes(spRoutes *gin.RouterGroup)
	// PostsRoutes(spRoutes *gin.RouterGroup)
	// PublicationsRoutes(spRoutes *gin.RouterGroup)
	// AnalyticRoute(spRoutes *gin.RouterGroup)
	// UserRoutes(spRoutes *gin.RouterGroup)
}
