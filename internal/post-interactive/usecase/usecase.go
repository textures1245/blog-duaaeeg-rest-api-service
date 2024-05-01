package usecase

import (
	_postInter "github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive/dtos"
	"github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive/entities"
)

type userInteractiveUse struct {
	CommRepo _postInter.CommentRepository
	LikeRepo _postInter.LikeRepository
}

func NewUserInteractiveUse(CommRepo _postInter.CommentRepository, LikeRepo _postInter.LikeRepository) _postInter.UserInteractiveService {
	return &userInteractiveUse{
		CommRepo,
		LikeRepo,
	}
}

func (u *userInteractiveUse) OnCreateNewComment(pUuid string, req *dtos.CommentReqDat) (*entities.CommentResDat, error) {
	comm, err := u.CommRepo.CreateComment(pUuid, req)
	if err != nil {
		return nil, err
	}
	return &entities.CommentResDat{
		UUID:      comm.UUID,
		UserUuid:  comm.UserUUID,
		Comment:   comm.Content,
		PostUUID:  comm.PostUUID,
		CreatedAt: comm.CreatedAt.String(),
		UpdateAt:  comm.UpdatedAt.String(),
	}, nil
}

func (u *userInteractiveUse) OnDeleteCommentByUUID(cUuid string) error {
	err := u.CommRepo.DeleteCommentByUUID(cUuid)
	if err != nil {
		return err
	}
	return nil
}

func (u *userInteractiveUse) OnUpdateCommentByUUID(cUuid string, req *dtos.CommentReqDat) (*entities.CommentResDat, error) {
	comm, err := u.CommRepo.UpdateCommentByUUID(cUuid, req)
	if err != nil {
		return nil, err
	}
	return &entities.CommentResDat{
		UUID:      comm.UUID,
		UserUuid:  comm.UserUUID,
		Comment:   comm.Content,
		PostUUID:  comm.PostUUID,
		CreatedAt: comm.CreatedAt.String(),
		UpdateAt:  comm.UpdatedAt.String(),
	}, nil
}

func (u *userInteractiveUse) OnLikedPost(pUuid string, req *dtos.LikeReqDat) (*entities.LikeResDat, error) {
	like, err := u.LikeRepo.CreateLike(pUuid, req)
	if err != nil {
		return nil, err
	}
	return &entities.LikeResDat{
		UUID:      like.UUID,
		UserUuid:  like.UserUUID,
		PostUuid:  like.PostUUID,
		CreatedAt: like.CreatedAt.String(),
	}, nil
}

func (u *userInteractiveUse) OnUnlikedPost(pUuid string, req *dtos.LikeReqDat) error {
	err := u.LikeRepo.DeleteLikeByUUID(pUuid, req.UserUuid)
	if err != nil {
		return err
	}
	return nil
}
