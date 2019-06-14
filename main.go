package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/api/routers"
	"github.com/phamvinhdat/project_news/api/routers/index"
	"github.com/phamvinhdat/project_news/database"
	"github.com/phamvinhdat/project_news/repository"
)

func setup(dbConfig *database.Config) *gin.Engine {
	//conect database
	conn, err := database.NewConnection(dbConfig)
	if err != nil { //handler database err
		println("Conect to database err")
	}

	//create repository
	userRepo := repository.NewMySQLUserRepo(conn)
	categoryRepo := repository.NewMySQLCategoryRepo(conn)

	//load static file
	r := gin.Default()
	r.Static("/css", "./public/view/css")
	r.Static("/images", "./public/view/images")
	r.Static("/js", "./public/view/js")
	r.LoadHTMLGlob("public/view/*.html")

	//create jwtauthen
	JWTAuthen := middleware.New()

	//create router
	routerIndex := index.New(userRepo, categoryRepo)
	router := routers.New(JWTAuthen, routerIndex)

	//create group
	groupIndex := r.Group("/")

	//regis router
	router.Register(groupIndex)

	return r
}

func main() {
	dbConfig := database.DefaultConfig()
	r := setup(dbConfig)

	r.Run()
}
