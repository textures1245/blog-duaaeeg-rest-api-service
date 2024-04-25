package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/api/db"
	"github.com/textures1245/BlogDuaaeeg-backend/api/routes/repository"
)

func InitRoute(spRoutes *gin.RouterGroup, db *db.PrismaClient) {
	r := repository.NewRouteRepository(db)

	r.RootRoutes(spRoutes)
	r.UserRoutes(spRoutes)
	r.PostsRoutes(spRoutes)
	// r.PublicationsRoutes(spRoutes)
	// r.AnalyticRoute(spRoutes)
}
