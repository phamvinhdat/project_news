package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/phamvinhdat/project_news/models"
)

type MySQLNewsTagRepo struct {
	Conn *gorm.DB
}

func NewMySQLNewsTagRepo(conn *gorm.DB) INewsTagRepo {
	return &MySQLNewsTagRepo{
		Conn: conn,
	}
}

func (t *MySQLNewsTagRepo) Create(newsTag *models.NewsTag) error {
	err := t.Conn.Debug().Create(newsTag).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *MySQLNewsTagRepo) FetchAllTagsOfNews(idNews int) ([]*models.Tag, error) {
	var tags []*models.Tag
	err := t.Conn.First(&tags, "news_id = ?", idNews).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}
