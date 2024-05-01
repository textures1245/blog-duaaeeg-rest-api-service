package postinteractive

import (
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive/entities"
)

type UserInteractiveService interface {
	OnCreateNewComment(postUuid string, req *dtos.CommentReqDat) (*entities.CommentResDat, error)
	OnUpdateCommentByUUID(cUuid string, req *dtos.CommentReqDat) (*entities.CommentResDat, error)
	OnDeleteCommentByUUID(cUuid string) error
	OnLikedPost(pUuid string, req *dtos.LikeReqDat) (*entities.LikeResDat, error)
	OnUnlikedPost(pUuid string, req *dtos.LikeReqDat) error
}
