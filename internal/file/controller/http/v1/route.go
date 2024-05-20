package v1

import (
	"github.com/gin-gonic/gin"
	_fileRepo "github.com/textures1245/BlogDuaaeeg-backend/internal/file/repository"
	_fileUsecase "github.com/textures1245/BlogDuaaeeg-backend/internal/file/usecase"
	_routeRepo "github.com/textures1245/BlogDuaaeeg-backend/pkg/datasource/route/repository"
)

type FileRepo struct {
	*_routeRepo.RouteRepo
}

func (routeRepo *FileRepo) FileRoutes(spRoutes *gin.RouterGroup) {
	fRg := spRoutes.Group("/file")

	fileRepo := _fileRepo.NewFileRepository(routeRepo.Db)
	fileUse := _fileUsecase.NewFileUsecase(fileRepo)
	fileConn := NewFileHandler(fileUse)

	fRg.POST("/upload", fileConn.UploadFile)
	fRg.GET("/get-files", fileConn.GetSourceFiles)
}
