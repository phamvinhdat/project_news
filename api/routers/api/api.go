package api

import (
	"log"
	"net/http"

	"github.com/phamvinhdat/project_news/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/models"
	"github.com/phamvinhdat/project_news/repository"
	"golang.org/x/crypto/bcrypt"
)

type RouterApi struct {
	UserRepo  repository.IUserRepo
	JwtAuthen *middleware.JWTAuthen
}

func NewRouterApi(userRepo repository.IUserRepo, jwtAuthen *middleware.JWTAuthen) *RouterApi {
	return &RouterApi{
		UserRepo:  userRepo,
		JwtAuthen: jwtAuthen,
	}
}

func (r *RouterApi) Register(group *gin.RouterGroup) {
	group.POST("/register", r.postRegister)
	group.POST("/login", r.postLogin)
}

func (r *RouterApi) postLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	err := r.JwtAuthen.Auth(username, password)
	if err != nil {
		log.Println("Loi nay vois username: ", username, err)
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	token := r.JwtAuthen.NewToken(username, 5)
	c.SetCookie("token", token, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "login success",
		"username": username,
	})
}

func (r *RouterApi) postRegister(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBind(&user)
	log.Println(user, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	err = r.JwtAuthen.Validate(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	user.Password = string(hashPassword)
	user.RoleID = 1 // default usernormal
	err = r.UserRepo.Create(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Register seccess",
	})
}
