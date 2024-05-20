package post

import (
	"github.com/gin-gonic/gin"
	_cateEntities "github.com/textures1245/BlogDuaaeeg-backend/internal/category/entities"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post/entities"
)

type PostService interface {
	OnCreateNewPost(c *gin.Context, cateResDat *_cateEntities.PostCategoryResDat, tagResDat *_cateEntities.PostTagResDat, req *dtos.PostReqDat) (*entities.PostResDat, error)
	OnFetchPostByUUID(uuid string) (*entities.PostResDat, error)
	OnFetchPublisherPosts(opts *dtos.FetchPostOptReq) ([]*entities.PostResDat, error)
	OnUpdatePostAndTagByUUID(c *gin.Context, cateResDat *_cateEntities.PostCategoryResDat, uuid string, req *dtos.PostReqDat) (*entities.PostResDat, error)
	OnSubmitPostToPublisher(userUuid string, postUuid string) (string, error)
	OnFetchOwnerPosts(userUuid string) ([]*entities.PostResDat, error)
	OnDeletePostByUUID(postUuid string) error
}
