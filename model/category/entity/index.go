package entity

type PostTagCateService interface {
	OnCreateOrUpdateCategory(postUuid string, req *PostCategoryReqDat) (*PostCategoryResDat, error)
	OnCreateTags(postUuid string, req *PostTagReqDat) (*PostTagResDat, error)
	OnUpdateTags(id int, req *PostTagReqDat) (*PostTagResDat, error)
	// OnUpdateCategory(cateId int, req *PostCategoryReqDat) (*PostCategoryResDat, error)
}
