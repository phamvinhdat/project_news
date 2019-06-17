package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/phamvinhdat/project_news/models"
)

type MySQLCategoryRepo struct {
	Conn *gorm.DB
}

func NewMySQLCategoryRepo(conn *gorm.DB) ICaregoryRepo {
	return &MySQLCategoryRepo{
		Conn: conn,
	}
}

func (c *MySQLCategoryRepo) FetchAll() ([]models.Category, error) {
	var categories []models.Category
	err := c.Conn.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *MySQLCategoryRepo) CountAll() int {
	var count int
	err := c.Conn.Model(&models.Category{}).Count(&count).Error
	if err != nil{
		return 0
	}

	return count
}
