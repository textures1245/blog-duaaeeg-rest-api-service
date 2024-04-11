package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/routes"
	"github.com/textures1245/BlogDuaaeeg-backend/utils"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":8080")
}

func runServer() {
	// setup
	r := gin.Default()
	lg := utils.NewConsoleLogger(utils.Level("TRACE"))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		lg.Db.Info("Defaulting to port %s", port)
	}

	// routes definition
	rG := r.Group("/api/v1")
	db := utils.DbConnect()
	routes.InitRoute(rG, db)

	lg.Db.Info("Listening on port %s", port)
	r.Run(":" + port)

}
