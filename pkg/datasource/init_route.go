package datasource

import (
	"github.com/gin-gonic/gin"
	"github.com/textures1245/BlogDuaaeeg-backend/db"
	_authR "github.com/textures1245/BlogDuaaeeg-backend/internal/auth/controller/http/v1"
	_postInterR "github.com/textures1245/BlogDuaaeeg-backend/internal/post-interactive/controller/http/v1"
	_postR "github.com/textures1245/BlogDuaaeeg-backend/internal/post/controller/http/v1"
	_userFollowerR "github.com/textures1245/BlogDuaaeeg-backend/internal/user-follower/controller/http/v1"
	_userR "github.com/textures1245/BlogDuaaeeg-backend/internal/user/controller/http/v1"
	"github.com/textures1245/BlogDuaaeeg-backend/pkg/datasource/route/repository"
)

func InitRoute(spRoutes *gin.RouterGroup, db *db.PrismaClient) {
	routeRepo := repository.RouteRepo{Db: db}
	ar := _authR.RouteRepo{RouteRepo: &routeRepo}
	ur := _userR.RouteRepo{RouteRepo: &routeRepo}
	usrFollowerR := _userFollowerR.RouteRepo{RouteRepo: &routeRepo}
	pr := _postR.RouteRepo{RouteRepo: &routeRepo}
	postInter := _postInterR.RouteRepo{RouteRepo: &routeRepo}

	ar.AuthRoutes(spRoutes)
	ur.UserRoutes(spRoutes)
	pr.PostsRoutes(spRoutes)
	usrFollowerR.UserFollowerRoutes(spRoutes)
	postInter.PostInteractiveRoutes(spRoutes)
	// r.PublicationsRoutes(spRoutes)
	// r.AnalyticRoute(spRoutes)
}
