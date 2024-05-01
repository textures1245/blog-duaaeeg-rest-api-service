package postinteractive

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive/dtos"
)

type LikeRepository interface {
	CreateLike(pUuid string, req *dtos.LikeReqDat) (*db.LikeModel, error)
	DeleteLikeByUUID(pUuid string, usrUuid string) error
}

type CommentRepository interface {
	CreateComment(postUuid string, req *dtos.CommentReqDat) (*db.CommentModel, error)
	UpdateCommentByUUID(cUuid string, req *dtos.CommentReqDat) (*db.CommentModel, error)
	DeleteCommentByUUID(cUuid string) error
}
