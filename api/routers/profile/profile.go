package profile

import (
	"log"
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
	cookie, _ := c.Request.Cookie("token")
	token := cookie.Value
	tk, _ := r.JwtAuthen.ParseToken(token)
	user, err := r.UserRepo.FetchByUsername(tk.Username)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	user.Password = ""
	user.RoleID = 0

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"status":   true,
		"userinfo": user,
	})
}

func (r *RouterProfile) getLogout(c *gin.Context) {
	tokenMock := r.JwtAuthen.NewToken("", 0)
	c.SetCookie("token", tokenMock, 7200, "/", "localhost", false, false)
	c.Redirect(http.StatusSeeOther, "/")
}
