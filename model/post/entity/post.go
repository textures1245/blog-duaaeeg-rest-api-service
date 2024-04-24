package entity

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	entityCate "github.com/textures1245/BlogDuaaeeg-backend/model/category/entity"
)

type PostReqDat struct {
	UserUuid     string                         `json:"user_uuid" db:"user_uuid" form:"user_uuid" binding:"required" validate:"required"`
	Title        string                         `json:"title" db:"title" form:"title" binding:"required" validate:"required"`
	Content      string                         `json:"content" db:"content" form:"content" binding:"required" validate:"required"`
	Published    bool                           `json:"published" db:"published" form:"published"`
	SrcType      string                         `json:"src_type" db:"src_type" form:"src_type" binding:"required" validate:"required"`
	PostCategory *entityCate.PostCategoryReqDat `json:"category" db:"category" form:"category" binding:"required"`
	PostTag      *entityCate.PostTagReqDat      `json:"tags" db:"tags" form:"tags" binding:"required"`
}

type PostResDat struct {
	UUID              string                         `json:"uuid"`
	UserUuid          string                         `json:"user_uuid"`
	Title             string                         `json:"title"`
	Source            string                         `json:"source"`
	Published         bool                           `json:"published"`
	SrcType           string                         `json:"src_type"`
	PublishedPostUUID string                         `json:"published_post_uuid"`
	PostUUID          string                         `json:"post_uuid"`
	Category          *entityCate.PostCategoryResDat `json:"category"`
	Tags              *entityCate.PostTagResDat      `json:"tags"`
	CreatedAt         string                         `json:"created_at"`
	UpdateAt          string                         `json:"update_at"`
}

type PostWithTagCateResDat struct {
	Post     *PostResDat                    `json:"post"`
	Category *entityCate.PostCategoryResDat `json:"category"`
	Tags     *entityCate.PostTagResDat      `json:"tags"`
}

type FetchPostOptReq struct {
	Page int `json:"page" form:"page" binding:"required" validate:"required"`
}

type PostRepository interface {
	CreatePost(cateResDat *entityCate.PostCategoryResDat, tagResDat *entityCate.PostTagResDat, req *PostReqDat) (*db.PostModel, error)
	FetchPostByUUID(uuid string) (*db.PostModel, error)
	FetchPublisherPosts(opts *FetchPostOptReq) ([]db.PublicationPostModel, error)
	FetchPostByUserUUID(userUuid string) ([]db.PostModel, error)
	UpdatePostByUUID(cateResDat *entityCate.PostCategoryResDat, uuid string, req *PostReqDat) (*db.PostModel, error)
	UpdatePostToPublisher(userUuid string, postUuid string) (string, error)
}

type PostService interface {
	OnCreateNewPost(cateResDat *entityCate.PostCategoryResDat, tagResDat *entityCate.PostTagResDat, req *PostReqDat) (*PostResDat, error)
	OnFetchPostByUUID(uuid string) (*PostResDat, error)
	OnFetchPublisherPosts(opts *FetchPostOptReq) ([]*PostResDat, error)
	OnUpdatePostAndTagByUUID(cateResDat *entityCate.PostCategoryResDat, uuid string, req *PostReqDat) (*PostResDat, error)
	OnSubmitPostToPublisher(userUuid string, postUuid string) (string, error)
	OnFetchOwnerPosts(userUuid string) ([]*PostResDat, error)
}

//TODO: add PostDelete
