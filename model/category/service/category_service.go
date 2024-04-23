package service

import "github.com/textures1245/BlogDuaaeeg-backend/model/category/entity"

type tagCateUse struct {
	cateRepo entity.PostCategoryRepository
	tagRepo  entity.PostTagRepository
}

func NewCategoryService(cateRepo entity.PostCategoryRepository, tagRepo entity.PostTagRepository) entity.PostTagCateService {
	return &tagCateUse{
		cateRepo,
		tagRepo,
	}
}

func (t *tagCateUse) OnCreateOrUpdateCategory(postUuid string, cateReq *entity.PostCategoryReqDat) (*entity.PostCategoryResDat, error) {
	categories, err := t.cateRepo.CreateOrUpdateCategory(postUuid, cateReq)
	if err != nil {
		return nil, err
	}

	res := &entity.PostCategoryResDat{
		ID:   categories.ID,
		Name: categories.Name,
	}

	return res, nil
}

// func (t *tagCateUse) OnUpdateCategory(cateId int, req *entity.PostCategoryReqDat) (*entity.PostCategoryResDat, error) {
// 	category, err := t.cateRepo.UpdateCategory(cateId, req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := &entity.PostCategoryResDat{
// 		ID:   category.ID,
// 		Name: category.Name,
// 	}

// 	return res, nil
// }

func (t *tagCateUse) OnCreateTags(postUuid string, tagReq *entity.PostTagReqDat) (*entity.PostTagResDat, error) {
	postTag, err := t.tagRepo.CreateTags(postUuid, tagReq)
	if err != nil {
		return nil, err
	}

	res := &entity.PostTagResDat{
		ID:   postTag.ID,
		Tags: postTag.Tags,
	}

	return res, nil
}

func (t *tagCateUse) OnUpdateTags(id int, req *entity.PostTagReqDat) (*entity.PostTagResDat, error) {
	postTag, err := t.tagRepo.UpdateTags(id, req)
	if err != nil {
		return nil, err
	}

	res := &entity.PostTagResDat{
		ID:   postTag.ID,
		Tags: postTag.Tags,
	}

	return res, nil
}
