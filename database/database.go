package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func getConnectionString(config *Config) string {
	var strMysql string
	strMysql = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.IP, config.Port, config.DatabaseName)
	return strMysql
}

// NewConnection return new connection to database
func NewConnection(conf *Config) (*gorm.DB, error) {
	db, err := gorm.Open(conf.Dialect, getConnectionString(conf))
	if err != nil {
		return nil, err
	}

	return db, nil
}
