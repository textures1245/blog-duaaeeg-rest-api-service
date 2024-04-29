package post

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type PostRepository interface {
	CreatePost(cateResDat *entityCate.PostCategoryResDat, tagResDat *entityCate.PostTagResDat, req *PostReqDat) (*db.PostModel, error)
	FetchPostByUUID(uuid string) (*db.PostModel, error)
	FetchPublisherPosts(opts *FetchPostOptReq) ([]db.PublicationPostModel, error)
	FetchPostByUserUUID(userUuid string) ([]db.PostModel, error)
	UpdatePostByUUID(cateResDat *entityCate.PostCategoryResDat, uuid string, req *PostReqDat) (*db.PostModel, error)
	UpdatePostToPublisher(userUuid string, postUuid string) (string, error)
	DeletePostByUUID(postUuid string) error
}
