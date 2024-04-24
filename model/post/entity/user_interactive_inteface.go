package entity

type UserInteractiveService interface {
	OnCreateNewComment(postUuid string, req *CommentReqDat) (*CommentResDat, error)
	OnUpdateCommentByUUID(cUuid string, req *CommentReqDat) (*CommentResDat, error)
	OnDeleteCommentByUUID(cUuid string) error
	OnLikedPost(pUuid string, req *LikeReqDat) (*LikeResDat, error)
	OnUnlikedPost(pUuid string, req *LikeReqDat) error
}
