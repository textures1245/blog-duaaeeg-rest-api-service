package repository

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	_routeEntity "github.com/textures1245/BlogDuaaeeg-backend/pkg/datasource/route"
)

type RouteRepo struct {
	Db *db.PrismaClient
}

func NewRouteRepository(db *db.PrismaClient) _routeEntity.RouteRepository {
	return &RouteRepo{
		Db: db,
	}
}
