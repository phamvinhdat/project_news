package admin

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/models"
	"github.com/phamvinhdat/project_news/repository"
)

type RouterAdmin struct {
	UserRepo     repository.IUserRepo
	JwtAuthen    *middleware.JWTAuthen
	CategoryRepo repository.ICaregoryRepo
	NewsRepo     repository.INewsRepo
	CensorRepo   repository.ICensorRepo
}

func NewRouterAdmin(userRepo repository.IUserRepo, jwtAuthen *middleware.JWTAuthen, categoryRepo repository.ICaregoryRepo, newsRepo repository.INewsRepo, censorRepo repository.ICensorRepo) *RouterAdmin {
	return &RouterAdmin{
		UserRepo:     userRepo,
		JwtAuthen:    jwtAuthen,
		CategoryRepo: categoryRepo,
		NewsRepo:     newsRepo,
		CensorRepo:   censorRepo,
	}
}

func (r *RouterAdmin) Register(group *gin.RouterGroup) {
	group.GET("/", r.getAdmin)
	group.GET("/news", r.getAdminNews)
	group.POST("/news/ispublic", r.postNewsIsPublic)
	group.DELETE("/news/:newsid", r.deleteNews)
}

func (r *RouterAdmin) deleteNews(c *gin.Context) {
	
}

func (r *RouterAdmin) postNewsIsPublic(c *gin.Context) {
	username := c.Request.Context().Value("username").(string)
	strIsPublic := c.PostForm("isPublic")
	strNewsID := c.PostForm("newsID")
	isPublic, err := strconv.ParseBool(strIsPublic)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	newsID, err := strconv.Atoi(strNewsID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	user, err := r.UserRepo.FetchByUsername(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	now := time.Now()
	censor, _ := r.CensorRepo.FetchByIDNews(newsID)
	if censor != nil {
		censor.DateCensor = &now
		censor.IsPublic = isPublic
	} else {
		censor = &models.Censor{
			UserID:     user.ID,
			NewsID:     newsID,
			IsPublic:   isPublic,
			DateCensor: &now,
			DatePublic: &now,
		}
		err = r.CensorRepo.Create(censor)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": err,
			})
			return
		}
	}

	err = r.CensorRepo.Update(censor)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Update successfully",
	})
}

func (r *RouterAdmin) getAdmin(c *gin.Context) {
	username := c.Request.Context().Value("username").(string)
	countCategory := r.CategoryRepo.CountAll()
	countNews := r.NewsRepo.CountAll()
	countUser := r.UserRepo.CountAll()

	c.HTML(http.StatusOK, "adminDashboard.html", gin.H{
		"status":   true,
		"username": username,
		"payload": gin.H{
			"countCategory": countCategory,
			"countNews":     countNews,
			"countUser":     countUser,
		},
	})
}

func (r *RouterAdmin) getAdminNews(c *gin.Context) {
	username := c.Request.Context().Value("username").(string)

	countCategory := r.CategoryRepo.CountAll()
	countNews := r.NewsRepo.CountAll()
	countUser := r.UserRepo.CountAll()
	news, _ := r.NewsRepo.FetchAllNew()

	var newsReturns []*newsReturn
	for _, val := range *news {
		censor, _ := r.CensorRepo.FetchByIDNews(val.ID)
		newsReturns = append(newsReturns, &newsReturn{News: val, Censor: censor})
	}

	c.HTML(http.StatusOK, "adminNews.html", gin.H{
		"status":   true,
		"username": username,
		"payload": gin.H{
			"countCategory": countCategory,
			"countNews":     countNews,
			"countUser":     countUser,
		},
		"news": newsReturns,
	})
}

type newsReturn struct {
	News   models.News
	Censor *models.Censor
}
