package repository

import (
	"github.com/phamvinhdat/project_news/models"
)

type IUserRepo interface {
	Create(user *models.User) error
	FetchByUsername(username string) (*models.User, error)
	FetchByEmail(email string) (*models.User, error)
}

type ICaregoryRepo interface{
	FetchAll()([]models.Category, error)
}

type IRoleRepo interface{

}

type ICommentRepo interface{
	Create(comment *models.Comment) error
	FetchByNews(idNews int)(*models.Comment, error)
}

type INewsRepo interface{
	Create(news *models.News) error
}
