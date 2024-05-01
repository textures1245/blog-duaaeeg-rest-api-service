package handler

import (
	"errors"
	"fmt"

	"github.com/textures1245/BlogDuaaeeg-backend/db"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
)

type HandleUse struct {
	HandleRepo entity.ErrorHandler
}

func NewHandler(errHandle entity.ErrorHandler) entity.ErrorHandler {
	return &HandleUse{
		HandleRepo: errHandle,
	}
}

func (h *HandleUse) PrismaAuthHandle(err entity.CError) *entity.CError {

	if errors.Is(err.Err, db.ErrNotFound) {
		err.Err = errors.New("UserNotFound")
		return &err
	} else if _, uniqueErr := db.IsErrUniqueConstraint(err.Err); uniqueErr {
		err.Err = errors.New("EmailAlreadyExists")
		return &err
	} else {
		return &err
	}
}

func (h *HandleUse) PrismaPostHandle(err entity.CError) *entity.CError {

	if errors.Is(err.Err, db.ErrNotFound) {
		err.Err = errors.New("PostNotFound")
		return &err
	} else {
		return &err
	}
}

func (h *HandleUse) PrismaCustomHandle(nameM string, err entity.CError) *entity.CError {

	if errors.Is(err.Err, db.ErrNotFound) {
		err.StatusCode = 404
		err.Err = errors.New(nameM + "NotFound")
		return &err
	} else if info, errUnique := db.IsErrUniqueConstraint(err.Err); errUnique {
		// Fields exists for Postgres and SQLite
		err.StatusCode = 409
		err.Err = fmt.Errorf("The action prevent for creating new unique constraint that already exist on the fields: %s", info.Fields)
		return &err
	} else {
		return &err
	}
}
