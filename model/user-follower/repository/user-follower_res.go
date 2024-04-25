package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	_errEntity "github.com/textures1245/BlogDuaaeeg-backend/error/entity"
	"github.com/textures1245/BlogDuaaeeg-backend/model/user-follower/entity"
)

type UsrFollowerRepo struct {
	db *db.PrismaClient
}

func NewUsrFollowerRepo(db *db.PrismaClient) entity.UserFollowerRepository {
	return &UsrFollowerRepo{db}
}

func (u *UsrFollowerRepo) CreateUserFollower(usrFollowerUuid string, req *entity.UserFollowerReqDat) (*db.UserFollowerModel, error) {

	ctx := context.Background()

	res, err := u.db.UserFollower.CreateOne(
		db.UserFollower.Followee.Link(
			db.User.UUID.Equals(req.UserFolloweeUuid),
		),
		db.UserFollower.Follower.Link(
			db.User.UUID.Equals(usrFollowerUuid),
		),
	).Exec(ctx)

	if err != nil {
		return nil, &_errEntity.CError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return res, nil
}

func (u *UsrFollowerRepo) DeleteUserFollowerByUUID(usrFollowerUuid string, req *entity.UserFollowerReqDat) error {
	ctx := context.Background()

	b, err := u.db.Prisma.ExecuteRaw(`DELETE FROM "UserFollower" WHERE "followerUuid" = $1 AND "followeeUuid" = $2`, usrFollowerUuid, req.UserFolloweeUuid).Exec(ctx)
	if err != nil {
		return &_errEntity.CError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	if b.Count == 0 {
		return &_errEntity.CError{
			Err:        errors.New("UserFollowerNotFoundOrAlreadyDeleted"),
			StatusCode: http.StatusNotFound,
		}
	}

	return nil
}
