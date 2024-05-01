package post

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	_cateEntities "github.com/textures1245/BlogDuaaeeg-backend/internal/category/entities"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post/dtos"
)

type PostRepository interface {
	CreatePost(cateResDat *_cateEntities.PostCategoryResDat, tagResDat *_cateEntities.PostTagResDat, req *dtos.PostReqDat) (*db.PostModel, error)
	FetchPostByUUID(uuid string) (*db.PostModel, error)
	FetchPublisherPosts(opts *dtos.FetchPostOptReq) ([]db.PublicationPostModel, error)
	FetchPostByUserUUID(userUuid string) ([]db.PostModel, error)
	UpdatePostByUUID(cateResDat *_cateEntities.PostCategoryResDat, uuid string, req *dtos.PostReqDat) (*db.PostModel, error)
	UpdatePostToPublisher(userUuid string, postUuid string) (string, error)
	DeletePostByUUID(postUuid string) error
}
