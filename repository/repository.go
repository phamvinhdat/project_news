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
