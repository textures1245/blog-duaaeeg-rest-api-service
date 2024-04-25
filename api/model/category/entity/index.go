package entity

type PostTagCateService interface {
	OnCreateOrUpdateCategory(req *PostCategoryReqDat) (*PostCategoryResDat, error)
	OnCreateTags(req *PostTagReqDat) (*PostTagResDat, error)
	OnUpdateTags(id int, req *PostTagReqDat) (*PostTagResDat, error)
	// OnUpdateCategory(cateId int, req *PostCategoryReqDat) (*PostCategoryResDat, error)
}
