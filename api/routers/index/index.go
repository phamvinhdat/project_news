package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/models"
	"github.com/phamvinhdat/project_news/repository"
)

type CategoryParent struct {
	ID            int
	Name          string
	CategoryChild []models.Category
}

type RouterIndex struct {
	UserRepo     repository.IUserRepo
	CategoryRepo repository.ICaregoryRepo
}

func New(userRepo repository.IUserRepo, categoryRepo repository.ICaregoryRepo) *RouterIndex {
	return &RouterIndex{
		UserRepo:     userRepo,
		CategoryRepo: categoryRepo,
	}
}

func (r *RouterIndex) Register(group *gin.RouterGroup) {
	group.GET("/", r.getIndex)
	group.GET("/login", r.getLogin)
	group.GET("/register", r.getRegister)
}

func (*RouterIndex) getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "signIn.html", gin.H{
		"title": "Đăng nhập",
	})
}

func (*RouterIndex) getRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Đăng kí",
	})
}

func (r *RouterIndex) getIndex(c *gin.Context) {
	categories, err := r.CategoryRepo.FetchAll()
	HandlerError(http.StatusNotFound, err, c)

	categoryParents := convertCategoriesToCategorytParents(categories)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":      "24 News — Tin tức 24h",
		"Categories": categoryParents,
	})
}

func convertCategoriesToCategorytParents(categories []models.Category) []CategoryParent {
	var categoryParents []CategoryParent
	for _, category := range categories {
		if category.ParentCategoryID > 0 {
			for i, _ := range categoryParents {
				if categoryParents[i].ID == category.ParentCategoryID {
					categoryParents[i].CategoryChild = append(categoryParents[i].CategoryChild, category)
					break
				}
			}
		} else {
			categoryParents = append(categoryParents, CategoryParent{
				ID:            category.ID,
				Name:          category.Name,
				CategoryChild: nil,
			})
		}
	}
	return categoryParents
}

func HandlerError(httpStatus int, err error, c *gin.Context) {
	if err != nil {
		c.AbortWithError(httpStatus, err)
		return
	}
}
