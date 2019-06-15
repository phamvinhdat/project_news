package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/api/routers/api"
	"github.com/phamvinhdat/project_news/api/routers/api/admin"
	"github.com/phamvinhdat/project_news/api/routers/api/writer"
	"github.com/phamvinhdat/project_news/api/routers/index"
	"github.com/phamvinhdat/project_news/api/routers/profile"
)

type Router struct {
	Middleware    *middleware.JWTAuthen
	RouterIndex   *index.RouterIndex
	RouterUsers   *api.RouterApi
	RouterAdmin   *admin.RouterAdmin
	RouterWriter  *writer.RouterWriter
	RouterProfile *profile.RouterProfile
}

func NewRouter(middleware *middleware.JWTAuthen, routerIndex *index.RouterIndex, routerApi *api.RouterApi, routerAdmin *admin.RouterAdmin, routerWriter *writer.RouterWriter, routerProfile *profile.RouterProfile) *Router {
	return &Router{
		Middleware:    middleware,
		RouterIndex:   routerIndex,
		RouterUsers:   routerApi,
		RouterAdmin:   routerAdmin,
		RouterWriter:  routerWriter,
		RouterProfile: routerProfile,
	}
}

func (r *Router) Register(groupIndex *gin.RouterGroup, groupApi *gin.RouterGroup, groupAdmin *gin.RouterGroup, groupWriter *gin.RouterGroup, groupProfile *gin.RouterGroup) {
	r.RouterIndex.Register(groupIndex)
	r.RouterUsers.Register(groupApi)
	r.RouterAdmin.Register(groupAdmin)
	r.RouterWriter.Register(groupWriter)
	r.RouterProfile.Register(groupProfile)
}
