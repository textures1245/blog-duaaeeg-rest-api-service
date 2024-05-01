package usecase

import (
	"github.com/textures1245/BlogDuaaeeg-backend/internal/category"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/category/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/category/entities"
)

type tagCateUse struct {
	cateRepo category.PostCategoryRepository
	tagRepo  category.PostTagRepository
}

func NewCategoryService(cateRepo category.PostCategoryRepository, tagRepo category.PostTagRepository) category.PostTagCateService {
	return &tagCateUse{
		cateRepo,
		tagRepo,
	}
}

func (t *tagCateUse) OnCreateOrUpdateCategory(cateReq *dtos.PostCategoryReqDat) (*entities.PostCategoryResDat, error) {
	categories, err := t.cateRepo.CreateOrUpdateCategory(cateReq)
	if err != nil {
		return nil, err
	}

	res := &entities.PostCategoryResDat{
		ID:   categories.ID,
		Name: categories.Name,
	}

	return res, nil
}

// func (t *tagCateUse) OnUpdateCategory(cateId int, req *category.PostCategoryReqDat) (*category.PostCategoryResDat, error) {
// 	category, err := t.cateRepo.UpdateCategory(cateId, req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := &category.PostCategoryResDat{
// 		ID:   category.ID,
// 		Name: category.Name,
// 	}

// 	return res, nil
// }

func (t *tagCateUse) OnCreateTags(tagReq *dtos.PostTagReqDat) (*entities.PostTagResDat, error) {
	postTag, err := t.tagRepo.CreateTags(tagReq)
	if err != nil {
		return nil, err
	}

	res := &entities.PostTagResDat{
		ID:   postTag.ID,
		Tags: postTag.Tags,
	}

	return res, nil
}

func (t *tagCateUse) OnUpdateTags(id int, req *dtos.PostTagReqDat) (*entities.PostTagResDat, error) {
	postTag, err := t.tagRepo.UpdateTags(id, req)
	if err != nil {
		return nil, err
	}

	res := &entities.PostTagResDat{
		ID:   postTag.ID,
		Tags: postTag.Tags,
	}

	return res, nil
}
