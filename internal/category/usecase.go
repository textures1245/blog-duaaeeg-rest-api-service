package category

import (
	"github.com/textures1245/BlogDuaaeeg-backend/internal/category/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/category/entities"
)

type PostTagCateService interface {
	OnCreateOrUpdateCategory(req *dtos.PostCategoryReqDat) (*entities.PostCategoryResDat, error)
	OnCreateTags(req *dtos.PostTagReqDat) (*entities.PostTagResDat, error)
	OnUpdateTags(id int, req *dtos.PostTagReqDat) (*entities.PostTagResDat, error)
	// OnUpdateCategory(cateId int, req *PostCategoryReqDat) (*PostCategoryResDat, error)
}
