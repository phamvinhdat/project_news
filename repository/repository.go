package repository

import (
	"github.com/phamvinhdat/project_news/models"
)

type IUserRepo interface {
	Create(user *models.User) error
	FetchByUsername(username string) (*models.User, error)
	FetchByEmail(email string) (*models.User, error)
	FetchRole(username string) (*models.Role, error)
	UpdatePassword(newPassword string, username string) error
	CountAll() int
	UpdatePhoneNumber(phoneNumber string,  username string) error
	UpdateName(newName string,  username string) error
}

type ICensorRepo interface {
	Create(*models.Censor) error
	FetchByIDNews(idNews int) (*models.Censor, error)
	Update(censor *models.Censor) error
}
type ITagRepo interface {
	Create(*models.Tag) error
	Fetch(id int) (*models.Tag, error)
	IsExists(name string) int
	FetchRandTag(number int) ([]models.Tag, error)
}

type INewsTagRepo interface {
	Create(newsTag *models.NewsTag) error
	FetchAllTagsOfNews(idNews int) ([]*models.Tag, error)
}

type ICaregoryRepo interface {
	FetchAll() ([]models.Category, error)
	CountAll() int
	FetchByName(name string) (*models.Category, error)
}

type IRoleRepo interface {
}

type ICommentRepo interface {
	Create(comment *models.Comment) error
	FetchByNews(idNews int) (*models.Comment, error)
}

type INewsRepo interface {
	CountAll() int
	Create(news *models.News) error
	FetchAllNew()(*[]models.News, error)
	FetchByID(newID int)(*models.News, error)
	FetchMostView(offset, number int, isPulic bool)([]models.News, error)
	FetchNewest(offset, number int, isPulic bool) ([]models.News, error)
	FetchRand(number int, isPulic bool) ([]models.News, error)
	FetchTopCategory(offset, number int, isPulic bool) ([]models.News, error)
	FetchNewestCategory(offset, number, categoryID, notEqualID int, isPulic bool) ([]models.News, error)
}
