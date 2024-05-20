package usecase

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/file"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/file/entities"
	errorEntity "github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
)

type fileUse struct {
	fileRepo file.FileRepository
}

func NewFileUsecase(fileRepo file.FileRepository) file.FileUsecase {
	return &fileUse{
		fileRepo,
	}
}

func (f *fileUse) GetSourceFiles(c *gin.Context, ctx context.Context) ([]*entities.File, int, error) {

	files, err := f.fileRepo.GetFiles(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, &errorEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return files, http.StatusOK, nil
}

func (f *fileUse) OnUploadFile(c *gin.Context, ctx context.Context, req *entities.FileUploaderReq) (*entities.File, int, error) {
	file := entities.File{
		FileName: req.FileName,
		FileType: req.FileType,
		FileData: req.FileData,
	}

	_, fPathDat, status, errOnDecode := file.EncodeBase64toFile(c, true)
	if errOnDecode != nil {
		return nil, status, errOnDecode
	}

	req.FileData = *fPathDat

	fileModel, err := f.fileRepo.CreateFile(ctx, req)
	if err != nil {
		return nil, http.StatusInternalServerError, &errorEntity.CError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return fileModel, http.StatusOK, nil
}
