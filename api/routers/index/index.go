package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
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
	JwtAuthen    *middleware.JWTAuthen
	NewsRepo     repository.INewsRepo
}

func NewRouterIndex(userRepo repository.IUserRepo, categoryRepo repository.ICaregoryRepo, jwtAuthen *middleware.JWTAuthen, newsRepo repository.INewsRepo) *RouterIndex {
	return &RouterIndex{
		UserRepo:     userRepo,
		CategoryRepo: categoryRepo,
		JwtAuthen:    jwtAuthen,
		NewsRepo:     newsRepo,
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
	//get category
	categories, err := r.CategoryRepo.FetchAll()
	HandlerError(http.StatusNotFound, err, c)
	categoryParents := convertCategoriesToCategorytParents(categories)

	//get cookie
	isLogin := false
	name := "Login"
	cookie, err := c.Request.Cookie("token")
	if err == nil {
		token := cookie.Value
		tk, err := r.JwtAuthen.ParseToken(token)
		if err == nil && tk.Username != "" {
			isLogin = true
			name = tk.Username
		}
	}

	mostViews, err := r.NewsRepo.FetchMostView(10, true)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	newest, err := r.NewsRepo.FetchNewest(10, true)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	randNews, err := r.NewsRepo.FetchRand(5, true)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	topCategoryNews, err := r.NewsRepo.FetchMostView(5, true)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":      "24 News — Tin tức 24h",
		"Categories": categoryParents,
		"isLogin":    isLogin,
		"name":       name,
		"news": gin.H{
			"mostViews": mostViews,
			"newest":    newest,
			"rand":      randNews,
			"top":       topCategoryNews,
		},
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

// func (r *RouterIndex)createTopCategoryAndRelate()([]topCategoryAndRelate), error){
	
// 	var result []topCategoryAndRelate
// 	topCategoryNews, err := r.NewsRepo.FetchTopCategory(10, true)

// 	 if err != nil{
// 		 return nil, err
// 	 }

// 	 for topNews := range topCategoryAndRelate{
// 		 relateNew, _ := r.newsRepo.FetchNewest(3, true)
// 		result = append(result, topCategoryAndRelate{
// 			News: topNews,
// 			RelateNew: relateNew,
// 		})
// 	 }

// 	 return result, nil
// }

type topCategoryAndRelate struct {
	News      *models.News
	RelateNew []*models.News
}
