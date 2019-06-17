package writer

import (
	"log"
	"net/http"
	"strings"
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
	NewsRepo        repository.INewsRepo
	TagRepo         repository.ITagRepo
	NewsTagRepo     repository.INewsTagRepo
}

func NewRouterWriter(userRepo repository.IUserRepo, jwtAuthen *middleware.JWTAuthen, category repository.ICaregoryRepo, imgLocalService services.IImg_service, newsRepo repository.INewsRepo, tagRepo repository.ITagRepo, newsTagRepo repository.INewsTagRepo) *RouterWriter {
	return &RouterWriter{
		UserRepo:        userRepo,
		JwtAuthen:       jwtAuthen,
		CategoryRepo:    category,
		ImgLocalService: imgLocalService,
		NewsRepo:        newsRepo,
		TagRepo:         tagRepo,
		NewsTagRepo:     newsTagRepo,
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
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	if role.Name != "writer" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

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

	user, err := r.UserRepo.FetchByUsername(tk.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Error when get userID",
		})
		return
	}

	datePost := time.Now()
	news.Avatar = strPath
	news.DatePost = &datePost
	news.UserID = user.ID
	err = r.NewsRepo.Create(&news)
	if err != nil {
		_ = r.ImgLocalService.Delete(strPath)
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Error when save news",
		})
		return
	}

	strTags := c.PostForm("tags")
	tags := strings.Split(strTags, ",")
	var tagID int
	for _, value := range tags {
		tagID = 0
		value = strings.TrimSpace(value)
		if value != "" {
			tag := models.Tag{Name: value}
			tagID = r.TagRepo.IsExists(value)
			if tagID == 0 {
				_ = r.TagRepo.Create(&tag)
				tagID = tag.ID
			}

		}

		if tagID != 0 {
			newsTag := models.NewsTag{NewsID: news.ID, TagID: tagID}
			log.Println(newsTag)
			_ = r.NewsTagRepo.Create(&newsTag)

		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"news":    news,
		"message": "Save news successfully",
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
