package file

import (
	"context"

	"github.com/textures1245/BlogDuaaeeg-backend/internal/file/entities"
)

type FileRepository interface {
	CreateFile(ctx context.Context, file *entities.FileUploaderReq) (*entities.File, error)
	GetFiles(ctx context.Context) ([]*entities.File, error)
	GetFileById(ctx context.Context, id *int64) (*entities.File, error)
}
