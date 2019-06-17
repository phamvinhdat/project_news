package main

import (
	"log"

	"github.com/phamvinhdat/project_news/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/phamvinhdat/project_news/api/middleware"
	"github.com/phamvinhdat/project_news/api/routers"
	"github.com/phamvinhdat/project_news/api/routers/api"
	"github.com/phamvinhdat/project_news/api/routers/api/admin"
	"github.com/phamvinhdat/project_news/api/routers/api/writer"
	"github.com/phamvinhdat/project_news/api/routers/index"
	"github.com/phamvinhdat/project_news/api/routers/profile"
	"github.com/phamvinhdat/project_news/database"
	"github.com/phamvinhdat/project_news/repository"
)

func setup(dbConfig *database.Config, conn *gorm.DB) *gin.Engine {

	//create repository
	userRepo := repository.NewMySQLUserRepo(conn)
	categoryRepo := repository.NewMySQLCategoryRepo(conn)
	newsRepo := repository.NewMySQLNewsRepo(conn)
	newsTagRepo := repository.NewMySQLNewsTagRepo(conn)
	tagRepo := repository.NewMySQLTagRepo(conn)

	//load static file
	r := gin.Default()
	r.Static("/css", "./public/view/css")
	r.Static("/images", "./public/view/images")
	r.Static("/js", "./public/view/js")
	r.LoadHTMLGlob("public/view/*.html")

	//create jwtauthen
	JWTAuthen := middleware.NewJWTAuthen(userRepo)

	//create service
	imgLocalService := services.NewImgLocalService()

	//create router
	routerIndex := index.NewRouterIndex(userRepo, categoryRepo, JWTAuthen)
	routerAPI := api.NewRouterApi(userRepo, JWTAuthen)
	routerAdmin := admin.NewRouterAdmin(userRepo, JWTAuthen)
	routerWriter := writer.NewRouterWriter(userRepo, JWTAuthen, categoryRepo, imgLocalService, newsRepo, tagRepo, newsTagRepo)
	routerProfile := profile.NewRouterProfile(userRepo, JWTAuthen)
	router := routers.NewRouter(JWTAuthen, routerIndex, routerAPI, routerAdmin, routerWriter, routerProfile)

	//create group
	groupIndex := r.Group("/")

	groupProfile := r.Group("/api/profile", JWTAuthen.JWTAuthentication())
	groupAPI := r.Group("/api")
	groupAdmin := r.Group("/api/admin", JWTAuthen.JWTAuthentication())
	groupWriter := r.Group("/api/writer", JWTAuthen.JWTAuthentication())

	//regis router
	router.Register(groupIndex, groupAPI, groupAdmin, groupWriter, groupProfile)

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
