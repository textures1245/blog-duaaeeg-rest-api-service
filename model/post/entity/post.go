package entity

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type PostReqDat struct {
	UserUuid  string `json:"user_uuid" db:"user_uuid" form:"user_uuid" binding:"required" validate:"required"`
	Title     string `json:"title" db:"title" form:"title" binding:"required" validate:"required"`
	Content   string `json:"content" db:"content" form:"content" binding:"required" validate:"required"`
	Published bool   `json:"published" db:"published" form:"published" binding:"required" validate:"required"`
	SrcType   string `json:"src_type" db:"src_type" form:"src_type" binding:"required" validate:"required"`
}

type PostResDat struct {
	UUID      string `json:"uuid"`
	UserUuid  string `json:"user_uuid"`
	Title     string `json:"title"`
	Source    string `json:"source"`
	Published bool   `json:"published"`
	SrcType   string `json:"src_type"`
}

type FetchPostOpts struct {
	Page int
}

type PostRepository interface {
	CreatePost(req *PostReqDat) (*db.PostModel, error)
	FetchPostByUUID(uuid string) (*db.PostModel, error)
	FetchPublisherPosts(opts *FetchPostOpts) ([]db.PublicationPostModel, error)
	FetchPostByUserUUID(userUuid string) ([]db.PostModel, error)
	UpdatePostByUUID(uuid string, req *PostReqDat) (*db.PostModel, error)
	UpdatePostToPublisher(userUuid string, postUuid string) error
}

type PostService interface {
	OnCreateNewPost(req *PostReqDat) (*PostResDat, error)
	OnFetchPostByUUID(uuid string) (*PostResDat, error)
	OnFetchPublisherPosts(opts *FetchPostOpts) ([]*PostResDat, error)
	OnUpdatePostByUUID(uuid string, req *PostReqDat) (*PostResDat, error)
	OnSubmitPostToPublisher(userUuid string, postUuid string) error
	OnFetchOwnerPosts(userUuid string) ([]*PostResDat, error)
}
