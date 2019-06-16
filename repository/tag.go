package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/phamvinhdat/project_news/models"
)

type MySQLTagRepo struct {
	Conn *gorm.DB
}

func NewMySQLTagRepo(conn *gorm.DB) ITagRepo {
	return &MySQLTagRepo{
		Conn: conn,
	}
}

func (t *MySQLTagRepo) Create(tag *models.Tag) error {
	err := t.Conn.Create(tag).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *MySQLTagRepo) IsExists(name string) int {
	var tag models.Tag
	err := t.Conn.First(&tag, "name = ?", name).Error
	if err != nil {
		return 0
	}

	log.Println(tag)
	return tag.ID
}

func (t *MySQLTagRepo) Fetch(id int) (*models.Tag, error) {
	var tag models.Tag
	err := t.Conn.First(&tag, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
