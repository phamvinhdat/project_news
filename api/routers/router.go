package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/api/routers/index"
)

type Router struct {
	Middleware  *middleware.JWTAuthen
	RouterIndex *index.RouterIndex
}

func New(middleware *middleware.JWTAuthen, routerIndex *index.RouterIndex) *Router {
	return &Router{
		Middleware:  middleware,
		RouterIndex: routerIndex,
	}
}

func (r *Router) Register(groupIndex *gin.RouterGroup) {
	r.RouterIndex.Register(groupIndex)
}
