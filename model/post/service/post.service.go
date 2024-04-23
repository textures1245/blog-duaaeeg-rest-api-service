package service

import (
	"fmt"
	"log"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	postEntity "github.com/textures1245/BlogDuaaeeg-backend/model/post/entity"
	userEntity "github.com/textures1245/BlogDuaaeeg-backend/model/user/entity"
)

type postUse struct {
	PostRepo  postEntity.PostRepository
	UsersRepo userEntity.UsersRepository
}

func NewPostService(postRepo postEntity.PostRepository, usersRepo userEntity.UsersRepository) postEntity.PostService {
	return &postUse{
		PostRepo:  postRepo,
		UsersRepo: usersRepo,
	}
}

func (u *postUse) OnCreateNewPost(cateId int, tagId int, req *postEntity.PostReqDat) (*postEntity.PostResDat, error) {
	post, err := u.PostRepo.CreatePost(cateId, tagId, req)
	if err != nil {
		return nil, err
	}

	// check if post marked as publish then link to publication post
	if _, isNil := post.PublicationPost(); !isNil {
		if post.Published {
			err := u.PostRepo.UpdatePostToPublisher(post.UserUUID, post.UUID)
			if err != nil {
				fmt.Println(fmt.Errorf("%v", err))
				return nil, err
			}
		}
	}

	var pbpUuid = ""
	if uuid, ok := post.PublishPostUUID(); ok {
		pbpUuid = uuid
	}

	res := &postEntity.PostResDat{
		UUID:              post.UUID,
		UserUuid:          post.UserUUID,
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: pbpUuid,
	}
	return res, nil

}

func (u *postUse) OnUpdatePostByUUID(cateId int, uuid string, req *postEntity.PostReqDat) (*postEntity.PostResDat, error) {
	post, err := u.PostRepo.UpdatePostByUUID(cateId, uuid, req)
	if err != nil {
		return nil, err
	}

	// check if post marked as publish then link to publication post
	if _, isNil := post.PublicationPost(); !isNil {
		log.Print(post)
		if post.Published {
			err := u.PostRepo.UpdatePostToPublisher(post.UserUUID, post.UUID)
			log.Println("error on update post to publisher", err)
			if err != nil {
				return nil, err
			}
		}
	}

	var pbpUuid = ""
	if uuid, ok := post.PublishPostUUID(); ok {
		pbpUuid = uuid
	}

	res := &postEntity.PostResDat{
		UUID:              post.UUID,
		UserUuid:          post.UserUUID,
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: pbpUuid,
	}
	return res, nil

}

func (u *postUse) OnFetchPostByUUID(uuid string) (*postEntity.PostResDat, error) {
	post, err := u.PostRepo.FetchPostByUUID(uuid)
	if err != nil {
		return nil, err
	}

	var pbpUuid = ""
	if uuid, ok := post.PublishPostUUID(); ok {
		pbpUuid = uuid
	}

	res := &postEntity.PostResDat{
		UUID:              post.UUID,
		UserUuid:          post.UserUUID,
		Title:             post.Title,
		Source:            post.Source,
		Published:         post.Published,
		SrcType:           string(post.SrcType),
		PublishedPostUUID: pbpUuid,
	}
	return res, nil
}

func (u *postUse) OnFetchOwnerPosts(userUuid string) ([]*postEntity.PostResDat, error) {
	posts, err := u.PostRepo.FetchPostByUserUUID(userUuid)
	if err != nil {
		return nil, err
	}

	var res []*postEntity.PostResDat
	res = mapPostsDatToRes(posts, res)
	return res, nil

}

func (u *postUse) OnFetchPublisherPosts(opts *postEntity.FetchPostOptReq) ([]*postEntity.PostResDat, error) {
	posts, err := u.PostRepo.FetchPublisherPosts(opts)
	if err != nil {
		return nil, err
	}

	var res []*postEntity.PostResDat
	for _, post := range posts {

		res = append(res, &postEntity.PostResDat{
			UUID:              post.UUID,
			UserUuid:          post.UserUUID,
			Title:             post.Post().Title,
			Source:            post.Post().Source,
			Published:         post.Post().Published,
			SrcType:           string(post.Post().SrcType),
			PublishedPostUUID: post.PostUUID,
		})
	}
	return res, nil
}

// - this func should be called as private func
func (u *postUse) OnSubmitPostToPublisher(userUuid string, postUuid string) error {
	err := u.PostRepo.UpdatePostToPublisher(userUuid, postUuid)
	if err != nil {
		return err
	}
	return nil
}

func mapPostsDatToRes(pDat []db.PostModel, pRes []*postEntity.PostResDat) []*postEntity.PostResDat {
	for _, post := range pDat {
		var pbpUuid = ""
		uuid, _ := post.PublishPostUUID()
		pbpUuid = uuid

		pRes = append(pRes, &postEntity.PostResDat{
			UUID:              post.UUID,
			UserUuid:          post.UserUUID,
			Title:             post.Title,
			Source:            post.Source,
			Published:         post.Published,
			SrcType:           string(post.SrcType),
			PublishedPostUUID: pbpUuid,
		})
	}
	return pRes
}
