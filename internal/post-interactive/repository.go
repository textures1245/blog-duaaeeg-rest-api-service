package postinteractive

import "github.com/textures1245/BlogDuaaeeg-backend/db"

type LikeRepository interface {
	CreateLike(pUuid string, req *LikeReqDat) (*db.LikeModel, error)
	DeleteLikeByUUID(pUuid string, usrUuid string) error
}

type CommentRepository interface {
	CreateComment(postUuid string, req *CommentReqDat) (*db.CommentModel, error)
	UpdateCommentByUUID(cUuid string, req *CommentReqDat) (*db.CommentModel, error)
	DeleteCommentByUUID(cUuid string) error
}
