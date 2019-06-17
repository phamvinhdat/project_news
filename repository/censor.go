package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/phamvinhdat/project_news/models"
)

type MySQLCenorRepo struct {
	Conn *gorm.DB
}

func NewMySQLCenorRepo(conn *gorm.DB) ICensorRepo {
	return &MySQLCenorRepo{
		Conn: conn,
	}
}

func (u *MySQLCenorRepo) Create(censor *models.Censor) error {
	err := u.Conn.Create(censor).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *MySQLCenorRepo) FetchByIDNews(idNews int) (*models.Censor, error) {
	var censor models.Censor
	err := u.Conn.First(&censor, "news_id = ?", idNews).Error
	if err != nil {
		return nil, err
	}

	return &censor, nil
}

func (u *MySQLCenorRepo) Update(censor *models.Censor) error {
	err := u.Conn.Save(censor).Error
	if err != nil {
		return nil
	}

	return err
}
