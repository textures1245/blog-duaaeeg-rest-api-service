package file

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/file/entities"
)

type FileUsecase interface {
	OnUploadFile(c *gin.Context, ctx context.Context, req *entities.FileUploaderReq) (*entities.File, int, error)
	GetSourceFiles(c *gin.Context, ctx context.Context) ([]*entities.File, int, error)
}
