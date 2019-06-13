package index

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/repository"
)

type RouterIndex struct {
	UserRepo *repository.IUserRepo
}

func New(userRepo *repository.IUserRepo) *RouterIndex {
	return &RouterIndex{
		UserRepo: userRepo,
	}
}

func (r *RouterIndex)Register(group *gin.RouterGroup){
	group.GET("/", getIndex)
}

func getIndex(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
