package entity

type PostTagCateService interface {
	OnCreateCategory(postUuid string, req *PostCategoryReqDat) (*PostCategoryResDat, error)
	OnUpdateCategory(cateId int, req *PostCategoryReqDat) (*PostCategoryResDat, error)
	OnCreateTags(postUuid string, req *PostTagReqDat) (*PostTagResDat, error)
	OnUpdateTags(id int, req *PostTagReqDat) (*PostTagResDat, error)
}
