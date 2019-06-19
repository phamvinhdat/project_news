package index

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"

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
	TagRepo      repository.ITagRepo
}

func NewRouterIndex(userRepo repository.IUserRepo, categoryRepo repository.ICaregoryRepo, jwtAuthen *middleware.JWTAuthen, newsRepo repository.INewsRepo, tagRepo repository.ITagRepo) *RouterIndex {
	return &RouterIndex{
		UserRepo:     userRepo,
		CategoryRepo: categoryRepo,
		JwtAuthen:    jwtAuthen,
		NewsRepo:     newsRepo,
		TagRepo:      tagRepo,
	}
}

func (r *RouterIndex) Register(group *gin.RouterGroup) {
	group.GET("/", r.getIndex)
	group.GET("/error", r.getError)
	group.GET("/login", r.getLogin)
	group.GET("/register", r.getRegister)
	group.GET("/posts/:category/:page", r.getCategory)
	group.GET("/post/:category/:newsID/:title", r.getPost)
}

func (r *RouterIndex) getCategory(c *gin.Context) {
	categoryName := c.Param("category")
	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)
	log.Println(categoryName)
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

	category, err := r.CategoryRepo.FetchByName(categoryName)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	news, err := r.NewsRepo.FetchNewestCategory(page*5, 5, category.ID, 0, true)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	if len(news) <= 0 {
		c.Redirect(http.StatusSeeOther, "/")
	}

	tags, err := r.TagRepo.FetchRandTag(10)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	randMews, err := r.NewsRepo.FetchRand(5, true)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	maxPage := 0
	mockNews, _ := r.NewsRepo.FetchNewestCategory(0, 150000, category.ID, 0, true)
	if mockNews != nil {
		maxPage = (len(mockNews) / 5) + 1
	}

	page0 := models.Page{Current: page - 2, IsSelect: false, CategoryName: categoryName, MaxPage: maxPage, NextPage: page - 1}
	page1 := models.Page{Current: page - 1, IsSelect: false, CategoryName: categoryName, MaxPage: maxPage, NextPage: page}
	page2 := models.Page{Current: page, IsSelect: true, CategoryName: categoryName, MaxPage: maxPage, NextPage: page + 1}
	page3 := models.Page{Current: page + 1, IsSelect: false, CategoryName: categoryName, MaxPage: maxPage, NextPage: page + 2}
	page4 := models.Page{Current: page + 2, IsSelect: false, CategoryName: categoryName, MaxPage: maxPage, NextPage: page + 3}

	pages := []models.Page{
		page0, page1, page2, page3, page4,
	}

	c.HTML(http.StatusOK, "category.html", gin.H{
		"Categories": categoryParents,
		"isLogin":    isLogin,
		"name":       name,
		"news":       news,
		"tags":       tags,
		"page":       pages,
		"randNews":   randMews,
		"prePage":    page - 1,
		"nextPage":   page + 1,
		"CategoryName": category.Name,
	})
}

func (*RouterIndex) getError(c *gin.Context) {

	c.HTML(http.StatusOK, "error.html", gin.H{})
}

func (r *RouterIndex) getPost(c *gin.Context) {
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

	strNewsID := c.Param("newsID")
	newsID, err := strconv.Atoi(strNewsID)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}
	news, err := r.NewsRepo.FetchByID(newsID)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	randTags, err := r.TagRepo.FetchRandTag(10)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	mostViews, err := r.NewsRepo.FetchMostView(0, 10, true)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	randMews, err := r.NewsRepo.FetchRand(10, true)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	c.HTML(http.StatusOK, "post.html", gin.H{
		"Categories": categoryParents,
		"isLogin":    isLogin,
		"name":       name,
		"news":       news,
		"content":    template.HTML(news.Content),
		"randTags":   randTags,
		"mostView":   mostViews,
		"randNews":   randMews,
	})
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

	mostViews, err := r.NewsRepo.FetchMostView(0, 10, true)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	newest, err := r.NewsRepo.FetchNewest(0, 10, true)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	randNews, err := r.NewsRepo.FetchRand(5, true)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	topCategoryNews, err := r.NewsRepo.FetchMostView(0, 5, true)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	topCategoryAndRelate, err := r.createTopCategoryAndRelate()
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	randMews2, err := r.NewsRepo.FetchRand(10, true)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	randTags, err := r.TagRepo.FetchRandTag(10)
	if err != nil {
		ctx := context.WithValue(c.Request.Context(), "error", err)
		c.Request = c.Request.WithContext(ctx)
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":      "24 News — Tin tức 24h",
		"Categories": categoryParents,
		"isLogin":    isLogin,
		"name":       name,
		"news": gin.H{
			"mostViews":    mostViews,
			"newest":       newest,
			"rand":         randNews,
			"top":          topCategoryNews,
			"topAndRelate": topCategoryAndRelate,
			"rand2":        randMews2,
			"randTag":      randTags,
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

func (r *RouterIndex) createTopCategoryAndRelate() ([]topCategoryAndRelate, error) {

	var result []topCategoryAndRelate
	topCategoryNews, err := r.NewsRepo.FetchTopCategory(0, 10, true)

	if err != nil {
		return nil, err
	}

	for _, topNews := range topCategoryNews {
		relateNew, _ := r.NewsRepo.FetchNewestCategory(0, 3, topNews.CategoryID, topNews.ID, true)
		result = append(result, topCategoryAndRelate{
			News:      topNews,
			RelateNew: relateNew,
		})
	}

	return result, nil
}

type topCategoryAndRelate struct {
	News      models.News
	RelateNew []models.News
}
