package handler

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/routes"
	"github.com/textures1245/BlogDuaaeeg-backend/utils"
)

func Handler(c *gin.Context) {
	// setup
	onProdMode := os.Getenv("GIN_MODE")

	var g *gin.Engine
	if onProdMode == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
		g = gin.Default()
	} else {
		g = gin.Default()
	}

	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = "8080"
	// 	lg.Db.Info("Defaulting to port %s", port, "") // Add a placeholder value as the final argument
	// }

	// routes definition
	rG := g.Group("/v1")
	db := utils.DbConnect()
	defer func() {
		if err := db.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	routes.InitRoute(rG, db)

	g.ServeHTTP(c.Writer, c.Request)

}
