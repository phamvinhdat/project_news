package writer

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/phamvinhdat/project_news/services"

	"github.com/phamvinhdat/project_news/models"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/repository"
)

type RouterWriter struct {
	UserRepo        repository.IUserRepo
	JwtAuthen       *middleware.JWTAuthen
	CategoryRepo    repository.ICaregoryRepo
	ImgLocalService services.IImg_service
}

func NewRouterWriter(userRepo repository.IUserRepo, jwtAuthen *middleware.JWTAuthen, category repository.ICaregoryRepo, imgLocalService services.IImg_service) *RouterWriter {
	return &RouterWriter{
		UserRepo:        userRepo,
		JwtAuthen:       jwtAuthen,
		CategoryRepo:    category,
		ImgLocalService: imgLocalService,
	}
}

func (r *RouterWriter) Register(group *gin.RouterGroup) {
	group.GET("/", r.getWriter)
	group.POST("/", r.postWriter)
}

func (r *RouterWriter) postWriter(c *gin.Context) {
	cookie, _ := c.Request.Cookie("token")
	token := cookie.Value
	tk, _ := r.JwtAuthen.ParseToken(token)
	role, err := r.UserRepo.FetchRole(tk.Username)
	log.Println("hihi", role, err)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	if role.Name != "writer" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	fmt.Println("hihi vao roi ne")

	var news models.News
	err = c.ShouldBind(&news)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"err":    err.Error(),
			"news":   news,
		})
		return
	}

	_, fileHeader, err := c.Request.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "file connot receive",
		})
		return
	}

	strPath, err := r.ImgLocalService.Save(fileHeader)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Error to save image to server",
		})
		return
	}

	datePost := time.Now()
	news.Avatar = strPath
	news.DatePost = &datePost

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"news":   news,
	})
}

func (r *RouterWriter) getWriter(c *gin.Context) {
	cookie, _ := c.Request.Cookie("token")
	token := cookie.Value
	tk, _ := r.JwtAuthen.ParseToken(token)
	role, err := r.UserRepo.FetchRole(tk.Username)
	log.Println("hihi", role, err)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	if role.Name != "writer" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	categories, err := r.CategoryRepo.FetchAll()
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	c.HTML(http.StatusOK, "addNews.html", gin.H{
		"status":     true,
		"name":       tk.Username,
		"categories": categories,
	})
}
