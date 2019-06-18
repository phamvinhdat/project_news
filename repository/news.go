package repository

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/phamvinhdat/project_news/models"
)

type MySQLNewsRepo struct {
	Conn *gorm.DB
}

func NewMySQLNewsRepo(conn *gorm.DB) *MySQLNewsRepo {
	return &MySQLNewsRepo{
		Conn: conn,
	}
}

func (n *MySQLNewsRepo) Create(news *models.News) error {
	err := n.Conn.Create(news).Error
	if err != nil {
		return err
	}

	return nil
}

func (n *MySQLNewsRepo) FetchMostView(offset, number int, isPulic bool) ([]models.News, error) {
	var news []models.News
	var err error

	if !isPulic {
		err = n.Conn.Offset(offset).Limit(number).Order("views desc").Find(&news).Error
	} else {
		err = n.Conn.Offset(offset).Limit(number).Joins("LEFT JOIN censors c ON news.id = c.news_id").Where("c.is_public = ?", isPulic).Order("views desc, date_post desc").Find(&news).Error
	}
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n *MySQLNewsRepo) FetchNewest(offset, number int, isPulic bool) ([]models.News, error) {
	var news []models.News
	var err error

	if !isPulic {
		err = n.Conn.Offset(offset).Limit(number).Order("date_post desc").Find(&news).Error
	} else {
		err = n.Conn.Offset(offset).Limit(number).Joins("LEFT JOIN censors c ON news.id = c.news_id").Where("c.is_public = ?", isPulic).Order("date_post desc").Find(&news).Error
	}
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n *MySQLNewsRepo) FetchNewestCategory(offset, number, categoryID, notEqualID int, isPulic bool) ([]models.News, error) {
	var news []models.News
	var err error

	if !isPulic {
		err = n.Conn.Offset(offset).Limit(number).Where("news.category_id = ? AND news.id <> ?", categoryID, notEqualID).Order("date_post desc").Find(&news).Error
	} else {
		err = n.Conn.Offset(offset).Debug().Limit(number).Joins("LEFT JOIN censors c ON news.id = c.news_id").Where("c.is_public = ? AND news.category_id = ? AND news.id <> ?", isPulic, categoryID, notEqualID).Order("date_post desc").Find(&news).Error
	}
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n *MySQLNewsRepo) FetchTopCategory(offset, number int, isPulic bool) ([]models.News, error) {
	var news []models.News
	var err error
	var strQuery string

	if !isPulic {
		strQuery = fmt.Sprintf("SELECT k.* FROM news k WHERE k.id = (SELECT n.id FROM news n WHERE n.category_id = k.category_id ORDER BY n.views DESC, date_post DESC LIMIT 1) ORDER BY k.views DESC, date_post DESC LIMIT %d OFFSET %d", number, offset)
	} else {
		strQuery = fmt.Sprintf("SELECT k.* FROM news k LEFT JOIN censors c on k.id = c.news_id WHERE c.is_public = %t AND k.id = (SELECT n.id FROM news n WHERE n.category_id = k.category_id ORDER BY n.views DESC, date_post DESC LIMIT 1) ORDER BY k.views DESC, date_post DESC LIMIT %d OFFSET %d", isPulic, number, offset)
	}
	rows, err := n.Conn.Raw(strQuery).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID         int
			Title      string
			Avatar     string
			Summary    string
			Content    string
			UserID     int
			DatePost   *time.Time
			CategoryID int
			Views      int
			IsPremium  bool
		)
		rows.Scan(&ID, &Title, &Avatar, &Summary, &Content, &UserID, &DatePost, &CategoryID, &Views, &IsPremium)
		article := models.News{
			ID:         ID,
			Title:      Title,
			Avatar:     Avatar,
			Summary:    Summary,
			Content:    Content,
			UserID:     UserID,
			DatePost:   DatePost,
			CategoryID: CategoryID,
			Views:      Views,
			IsPremium:  IsPremium,
		}
		println(ID)
		news = append(news, article)
	}

	return news, nil
}

func (n *MySQLNewsRepo) FetchRand(number int, isPulic bool) ([]models.News, error) {
	var news []models.News
	var err error

	if !isPulic {
		err = n.Conn.Debug().Limit(number).Order("RAND()").Find(&news).Error
	} else {
		err = n.Conn.Debug().Limit(number).Joins("LEFT JOIN censors c ON news.id = c.news_id").Where("c.is_public = ?", isPulic).Order("RAND()").Find(&news).Error
	}
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n *MySQLNewsRepo) FetchAllNew() (*[]models.News, error) {
	var news []models.News
	err := n.Conn.Order("date_post DESC").Find(&news).Error
	if err != nil {
		return nil, err
	}

	return &news, nil
}

func (n *MySQLNewsRepo) FetchByID(newID int) (*models.News, error) {
	var news models.News
	err := n.Conn.First(&news, "id = ?", newID).Error
	if err != nil {
		return nil, err
	}

	return &news, nil
}

func (n *MySQLNewsRepo) CountAll() int {
	var count int
	err := n.Conn.Model(&models.News{}).Count(&count).Error
	if err != nil {
		return 0
	}

	return count
}
func (n *MySQLNewsRepo) PageByNews(limit int, offset int) (*[]models.News, error) {
	var news []models.News

	err := n.Conn.Debug().Offset(offset).Limit(limit).Find(&news).Error
	if err != nil {
		return nil, err
	}

	return &news, nil
}
