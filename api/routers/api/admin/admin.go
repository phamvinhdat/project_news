package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/repository"
)

type RouterAdmin struct {
	UserRepo     repository.IUserRepo
	JwtAuthen    *middleware.JWTAuthen
	CategoryRepo repository.ICaregoryRepo
	NewsRepo     repository.INewsRepo
}

func NewRouterAdmin(userRepo repository.IUserRepo, jwtAuthen *middleware.JWTAuthen, categoryRepo repository.ICaregoryRepo, newsRepo repository.INewsRepo) *RouterAdmin {
	return &RouterAdmin{
		UserRepo:     userRepo,
		JwtAuthen:    jwtAuthen,
		CategoryRepo: categoryRepo,
		NewsRepo:     newsRepo,
	}
}

func (r *RouterAdmin) Register(group *gin.RouterGroup) {
	group.GET("/", r.getAdmin)
	group.GET("/news", r.getAdminNews)
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

	c.HTML(http.StatusOK, "adminNews.html", gin.H{
		"status":   true,
		"username": username,
		"payload": gin.H{
			"countCategory": countCategory,
			"countNews":     countNews,
			"countUser":     countUser,
		},
		"news": news,
	})
}
