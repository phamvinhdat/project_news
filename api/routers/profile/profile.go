package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/repository"
	"golang.org/x/crypto/bcrypt"
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
	group.POST("/password", r.postPassword)
	// group.GET("/:name/:phone", r.changeInfo)
}



func (r *RouterProfile) postPassword(c *gin.Context) {
	cookie, _ := c.Request.Cookie("token")
	token := cookie.Value

	tk, _ := r.JwtAuthen.ParseToken(token)
	oldPassword := c.PostForm("oldPass")
	findUser, err := r.UserRepo.FetchByUsername(tk.Username)
	if err != nil {
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"status":   false,
			"userinfo": err,
		})
		return
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(oldPassword))
	if err != nil {
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"status":   false,
			"userinfo": err,
		})
		return
	}

	newPassword := c.PostForm("newPass")
	rePassword := c.PostForm("rePass")
	if newPassword != rePassword{
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"status":   false,
			"userinfo": "Repeat password incorect.",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"status":   false,
			"userinfo": err,
		})
		return
	}

	err = r.UserRepo.UpdatePassword(string(hashedPassword), tk.Username)
	if err != nil {
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"status":   false,
			"userinfo": err,
		})
		return
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"status":   true,
	})
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
