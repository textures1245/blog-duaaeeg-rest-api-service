package service

import (
	"github.com/textures1245/BlogDuaaeeg-backend/api/model/post/entity"
)

type userInteractiveUse struct {
	CommRepo entity.CommentRepository
	LikeRepo entity.LikeRepository
}

func NewUserInteractiveUse(CommRepo entity.CommentRepository, LikeRepo entity.LikeRepository) entity.UserInteractiveService {
	return &userInteractiveUse{
		CommRepo,
		LikeRepo,
	}
}

func (u *userInteractiveUse) OnCreateNewComment(pUuid string, req *entity.CommentReqDat) (*entity.CommentResDat, error) {
	comm, err := u.CommRepo.CreateComment(pUuid, req)
	if err != nil {
		return nil, err
	}
	return &entity.CommentResDat{
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

func (u *userInteractiveUse) OnUpdateCommentByUUID(cUuid string, req *entity.CommentReqDat) (*entity.CommentResDat, error) {
	comm, err := u.CommRepo.UpdateCommentByUUID(cUuid, req)
	if err != nil {
		return nil, err
	}
	return &entity.CommentResDat{
		UUID:      comm.UUID,
		UserUuid:  comm.UserUUID,
		Comment:   comm.Content,
		PostUUID:  comm.PostUUID,
		CreatedAt: comm.CreatedAt.String(),
		UpdateAt:  comm.UpdatedAt.String(),
	}, nil
}

func (u *userInteractiveUse) OnLikedPost(pUuid string, req *entity.LikeReqDat) (*entity.LikeResDat, error) {
	like, err := u.LikeRepo.CreateLike(pUuid, req)
	if err != nil {
		return nil, err
	}
	return &entity.LikeResDat{
		UUID:      like.UUID,
		UserUuid:  like.UserUUID,
		PostUuid:  like.PostUUID,
		CreatedAt: like.CreatedAt.String(),
	}, nil
}

func (u *userInteractiveUse) OnUnlikedPost(pUuid string, req *entity.LikeReqDat) error {
	err := u.LikeRepo.DeleteLikeByUUID(pUuid, req.UserUuid)
	if err != nil {
		return err
	}
	return nil
}
