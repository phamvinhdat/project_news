package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/api/routers/index"
	"github.com/phamvinhdat/project_news/api/routers/api"
)

type Router struct {
	Middleware  *middleware.JWTAuthen
	RouterIndex *index.RouterIndex
	RouterUsers *api.RouterApi
}

func NewRouter(middleware *middleware.JWTAuthen, routerIndex *index.RouterIndex, routerApi *api.RouterApi) *Router {
	return &Router{
		Middleware:  middleware,
		RouterIndex: routerIndex,
		RouterUsers: routerApi,
	}
}

func (r *Router) Register(groupIndex *gin.RouterGroup, groupApi *gin.RouterGroup) {
	r.RouterIndex.Register(groupIndex)
	r.RouterUsers.Register(groupApi)
}
