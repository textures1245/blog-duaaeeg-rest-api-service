package repository

import (
	"context"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/file"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/file/entities"
)

type fileRepo struct {
	db *db.PrismaClient
}

func NewFileRepository(db *db.PrismaClient) file.FileRepository {
	return &fileRepo{
		db: db,
	}
}

func (r *fileRepo) GetFiles(ctx context.Context) ([]*entities.File, error) {
	fileM, err := r.db.File.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}

	files := make([]*entities.File, 0)
	for _, f := range fileM {
		files = append(files, &entities.File{
			Id:        f.ID,
			FileName:  f.FileName,
			FileType:  f.FileType,
			FileData:  f.FileURL,
			CreatedAt: f.CreatedAt.Local().String(),
			UpdatedAt: f.UpdatedAt.Local().String(),
		})
	}

	return files, nil
}

func (r *fileRepo) CreateFile(ctx context.Context, file *entities.FileUploaderReq) (*entities.File, error) {
	fileModel, err := r.db.File.CreateOne(
		db.File.FileName.Set(file.FileName),
		db.File.FileType.Set(file.FileType),
		db.File.FileURL.Set(file.FileData),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	f := &entities.File{
		Id:        fileModel.ID,
		FileName:  fileModel.FileName,
		FileType:  fileModel.FileType,
		FileData:  fileModel.FileURL,
		CreatedAt: fileModel.CreatedAt.Local().String(),
		UpdatedAt: fileModel.UpdatedAt.Local().String(),
	}

	return f, nil

}

func (r *fileRepo) GetFileById(ctx context.Context, id *int64) (*entities.File, error) {
	fileModel, err := r.db.File.FindUnique(db.File.ID.Equals(int(*id))).Exec(ctx)
	if err != nil {
		return nil, err
	}

	f := &entities.File{
		Id:        fileModel.ID,
		FileName:  fileModel.FileName,
		FileType:  fileModel.FileType,
		FileData:  fileModel.FileURL,
		CreatedAt: fileModel.CreatedAt.Local().String(),
		UpdatedAt: fileModel.UpdatedAt.Local().String(),
	}

	return f, nil
}
