package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/repository"
)

type RouterAdmin struct {
	UserRepo  repository.IUserRepo
	JwtAuthen *middleware.JWTAuthen
}

func NewRouterAdmin(userRepo repository.IUserRepo, jwtAuthen *middleware.JWTAuthen) *RouterAdmin {
	return &RouterAdmin{
		UserRepo:  userRepo,
		JwtAuthen: jwtAuthen,
	}
}

func (r *RouterAdmin) Register(group *gin.RouterGroup) {
	group.GET("/", r.getAdmin)
}

func (r *RouterAdmin) getAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "you are admin",
	})
}
