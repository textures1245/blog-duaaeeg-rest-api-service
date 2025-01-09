//- This where entry point of the application

package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	// Configure logging
	logLevel := os.Getenv("LOG_LEVEL")

	logConfig := &utils.Logger{
		LogLevel: logLevel,
		LogFormat: &formatter.Formatter{
			CallerFirst: true,
			CustomCallerFormatter: func(f *runtime.Frame) string {
				s := strings.Split(f.Function, ".")
				funcName := s[len(s)-1]
				return fmt.Sprintf(" [%s:%d][%s()]", path.Base(f.File), f.Line, funcName)
			},
		},
		LogFilePath: os.Getenv("LOG_FILE_PATH"),
	}
	logConfig.InitConfig()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Infof("Defaulting to port %s", port) // Add a placeholder value as the final argument
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
	r.Static("/public/image", "./public/image")

	log.Infof("Listening on port %s", port)
	r.Run(":" + port)

}
