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
	err := u.Conn.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *MySQLUserRepo) FetchByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.Conn.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *MySQLUserRepo) FetchRole(username string) (*models.Role, error) {
	user, err := u.FetchByUsername(username)
	if err != nil {
		return nil, err
	}

	var role models.Role
	err = u.Conn.First(&role, "id = ?", user.RoleID).Error
	if err != nil {
		return nil, err
	}

	return &role, err
}

func (u *MySQLUserRepo) UpdatePassword(newPassword string,  username string) error {
	user, err := u.FetchByUsername(username)
	if err != nil {
		return err
	}

	err = u.Conn.Model(user).Update("password", newPassword).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *MySQLUserRepo) CountAll() int {
	var count int
	err := u.Conn.Model(&models.User{}).Count(&count).Error
	if err != nil {
		return 0
	}

	return count
}
