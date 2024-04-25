package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/routes"
	"github.com/textures1245/BlogDuaaeeg-backend/utils"
)

//go:generate go run github.com/steebchen/prisma-client-go generate

func Handler(w http.ResponseWriter, r *http.Request) {
	// setup
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	router := gin.Default()

	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = "8080"
	// 	lg.Db.Info("Defaulting to port %s", port, "") // Add a placeholder value as the final argument
	// }

	// routes definition
	rG := router.Group("/v1")
	db := utils.DbConnect()
	defer func() {
		if err := db.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	routes.InitRoute(rG, db)

	router.ServeHTTP(w, r)
}
