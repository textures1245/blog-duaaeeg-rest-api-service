package repository

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
)

type RouteRepo struct {
	Db *db.PrismaClient
}

//- TEMP: commented code for temp fix
// func NewRouteRepository(db *db.PrismaClient) _routeEntity.RouteRepository {
// 	return &RouteRepo{
// 		Db: db,
// 	}
// }
