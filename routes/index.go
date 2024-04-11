package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/routes/repository"
)

func InitRoute(spRoutes *gin.RouterGroup, db *db.PrismaClient) {
	r := repository.NewRouteRepository(db)

	r.RootRoutes(spRoutes)
	// r.PostsRoutes(spRoutes)
	// r.PublicationsRoutes(spRoutes)
	// r.AnalyticRoute(spRoutes)
	// r.UserRoutes(spRoutes)
}
