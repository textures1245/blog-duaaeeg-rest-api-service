package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/file"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/file/entities"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/error/handler"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/utils"
	// "github.com/textures1245/go-template/internal/file/usecase"
)

type fileConn struct {
	FileUse file.FileUsecase
}

func NewFileHandler(authUse file.FileUsecase) *fileConn {
	return &fileConn{
		FileUse: authUse,
	}
}

func (h *fileConn) UploadFile(c *gin.Context) {
	req := new(entities.FileUploaderReq)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	errOnValidate := utils.SchemaValidator(req)
	if errOnValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     fmt.Sprintf("error, invalid validated on schema body: %v", errOnValidate),
			"result":      nil,
		})
		return
	}

	var (
		ctx, cancel = context.WithTimeout(c.Request.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	res, status, err := h.FileUse.OnUploadFile(c, ctx, req)
	if err != nil {
		handlerE := handler.NewHandler(&handler.HandleUse{})
		hE := handlerE.PrismaPostHandle(*err.(*_errEntity.CError))
		c.JSON(status, gin.H{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     hE.Error(),
			"result":      nil,
		})
		return
	}

	if res.FileType == "PDF" {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", res.FileName))
		c.File(res.FileData)
		return
	}
	c.JSON(status, gin.H{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      res,
	})
	return
}

func (h *fileConn) GetSourceFiles(c *gin.Context) {
	var (
		ctx, cancel = context.WithTimeout(c.Request.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	files, status, err := h.FileUse.GetSourceFiles(c, ctx)
	if err != nil {
		handlerE := handler.NewHandler(&handler.HandleUse{})
		hE := handlerE.PrismaPostHandle(*err.(*_errEntity.CError))
		c.JSON(status, gin.H{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     hE.Error(),
			"result":      nil,
		})
		return
	}

	c.JSON(status, gin.H{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      files,
	})
	return
}
