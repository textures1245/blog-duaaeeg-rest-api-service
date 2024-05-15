//- This where entry point of the application

package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/datasource"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/utils"
)

func main() {
	// setup
	onProdMode := os.Getenv("GIN_MODE")

	var r *gin.Engine
	if onProdMode == "release" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.Default()
	} else {
		r = gin.Default()
	}
	lg := utils.NewConsoleLogger(utils.Level("TRACE"))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		lg.Db.Info("Defaulting to port %s", port, "") // Add a placeholder value as the final argument
	}

	// routes definition
	rG := r.Group("/api/v1")
	db := datasource.DbConnect()
	defer func() {
		if err := db.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	datasource.InitRoute(rG, db)

	lg.Db.Info("Listening on port %s", port, "")
	r.Run(":" + port)

}
