package main

import (
	"github.com/jinzhu/gorm"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/api/routers"
	"github.com/phamvinhdat/project_news/api/routers/api"
	"github.com/phamvinhdat/project_news/api/routers/index"
	"github.com/phamvinhdat/project_news/database"
	"github.com/phamvinhdat/project_news/repository"
)

func setup(dbConfig *database.Config, conn *gorm.DB) *gin.Engine {

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
	JWTAuthen := middleware.NewJWTAuthen(userRepo)

	//create router
	routerIndex := index.NewRouterIndex(userRepo, categoryRepo)
	routerApi := api.NewRouterApi(userRepo, JWTAuthen)
	router := routers.NewRouter(JWTAuthen, routerIndex, routerApi)

	//create group
	groupIndex := r.Group("/")
	groupApi := r.Group("/api")

	//regis router
	router.Register(groupIndex, groupApi)

	return r
}

func main() {
	//conect database
	dbConfig := database.DefaultConfig()
	conn, err := database.NewConnection(dbConfig)
	if err != nil { //handler database err
		log.Fatal("database err:", err)
	}
	defer conn.Close()
	r := setup(dbConfig, conn)

	r.Run()
}
