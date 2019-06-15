package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/repository"
)

type RouterProfile struct {
	UserRepo  repository.IUserRepo
	JwtAuthen *middleware.JWTAuthen
}

func NewRouterProfile(userRepo repository.IUserRepo, jwtAuthen *middleware.JWTAuthen) *RouterProfile {
	return &RouterProfile{
		UserRepo:  userRepo,
		JwtAuthen: jwtAuthen,
	}
}

func (r *RouterProfile) Register(group *gin.RouterGroup) {
	group.GET("/", r.getProfile)
	group.GET("/logout", r.getLogout)
}

func (r *RouterProfile) getProfile(c *gin.Context) {

}

func (r *RouterProfile) getLogout(c *gin.Context) {
	tokenMock := r.JwtAuthen.NewToken("Login", 0)
	c.SetCookie("token", tokenMock, 7200, "/", "localhost", false, false)
	c.Redirect(http.StatusSeeOther, "/")
}
