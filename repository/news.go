package repository

import (
	"github.com/phamvinhdat/project_news/models"
	"github.com/jinzhu/gorm"
)

type MySQLNewsRepo struct {
	Conn *gorm.DB
}

func NewMySQLNewsRepo(conn *gorm.DB) *MySQLNewsRepo{
	return &MySQLNewsRepo{
		Conn: conn,
	}
}

func (n *MySQLNewsRepo)Create(news *models.News)error{
	err := n.Conn.Create(news).Error
	if err != nil{
		return err
	}

	return nil
}

func (n *MySQLNewsRepo)FetchAllNew()(*[]models.News, error){
	var news []models.News
	err := n.Conn.Find(&news).Error
	if err != nil{
		return nil, err
	}

	return &news, nil
}

func (n *MySQLNewsRepo)FetchByID(newID int)(*models.News, error){
	var news models.News
	err := n.Conn.First(&news, "id = ?", newID).Error
	if err != nil{
		return nil, err
	}

	return &news, nil
}

func (n *MySQLNewsRepo) CountAll() int {
	var count int
	err := n.Conn.Model(&models.News{}).Count(&count).Error
	if err != nil{
		return 0
	}

	return count
}
func (n *MySQLNewsRepo)PageByNews(limit int, offset int)(*[]models.News, error){
	var news []models.News

	err := n.Conn.Debug().Offset(offset).Limit(limit).Find(&news).Error
	if err != nil{
		return nil, err
	}

	return &news, nil
}