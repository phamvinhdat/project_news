package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/phamvinhdat/project_news/models"
)

type MySQLUserRepo struct {
	Conn *gorm.DB
}

func NewMySQLUserRepo(conn *gorm.DB) IUserRepo {
	return &MySQLUserRepo{
		Conn: conn,
	}
}

func (u *MySQLUserRepo) Create(user *models.User) error {
	err := u.Conn.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *MySQLUserRepo) FetchByUsername(username string) (*models.User, error) {
	var user models.User
	err := u.Conn.First(user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *MySQLUserRepo) FetchByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.Conn.First(user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
