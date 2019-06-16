package writer

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/repository"
)

type RouterWriter struct {
	UserRepo     repository.IUserRepo
	JwtAuthen    *middleware.JWTAuthen
	CategoryRepo repository.ICaregoryRepo
}

func NewRouterWriter(userRepo repository.IUserRepo, jwtAuthen *middleware.JWTAuthen, category repository.ICaregoryRepo) *RouterWriter {
	return &RouterWriter{
		UserRepo:     userRepo,
		JwtAuthen:    jwtAuthen,
		CategoryRepo: category,
	}
}

func (r *RouterWriter) Register(group *gin.RouterGroup) {
	group.GET("/", r.getWriter)
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
