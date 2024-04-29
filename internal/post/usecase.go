package post

type PostService interface {
	OnCreateNewPost(cateResDat *entityCate.PostCategoryResDat, tagResDat *entityCate.PostTagResDat, req *PostReqDat) (*PostResDat, error)
	OnFetchPostByUUID(uuid string) (*PostResDat, error)
	OnFetchPublisherPosts(opts *FetchPostOptReq) ([]*PostResDat, error)
	OnUpdatePostAndTagByUUID(cateResDat *entityCate.PostCategoryResDat, uuid string, req *PostReqDat) (*PostResDat, error)
	OnSubmitPostToPublisher(userUuid string, postUuid string) (string, error)
	OnFetchOwnerPosts(userUuid string) ([]*PostResDat, error)
	OnDeletePostByUUID(postUuid string) error
}
