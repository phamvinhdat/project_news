package repository

import (
	"time"

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

func (t *MySQLTagRepo) FetchTopTag(number int, timeDuring *time.Time, isPublic, isPrivate bool) ([]models.Tag, error) {
	// var tags []models.Tag
	
	// if !isPublic{
	// 	err := t.Conn.Debug().Limit(number)..Error
	// }else{

	// }

	// if err != nil {
	// 	return nil, err
	// }

	// return tags, nil
	return nil, nil
}

func (t *MySQLTagRepo) FetchRandTag(number int) ([]models.Tag, error) {
	var tags []models.Tag
	err := t.Conn.Debug().Limit(number).Order("RAND()").Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *MySQLTagRepo) FetchByName(name string) (*models.Tag, error) {
	var tag models.Tag
	err := t.Conn.First("name = ?", name).Find(&tag).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
