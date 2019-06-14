package models

import (
	"time"
)

type News struct {
	ID         int        `gorm:"primary_key;auto_increment"`
	Title      string     `gorm:"column:title;type:nvarchar(255);not null"`
	Avatar     string     `gorm:"column:avatar;type:nvarchar(255); not null"`
	Summary    string     `gorm:"column:summary;type:text;not null"`
	Content    string     `gorm:"column:content;type:longtext;not null"`
	UserID     int        `gorm:"column:user_id;type:int(11) not null"`
	DatePost   *time.Time `gorm:"column:date_post;type:datetime not null"`
	CategoryID int        `gorm:"column:category_id;type:int(11) not null"`
	Views      int        `gorm:"colums:views;type:int(11);default 0"`
}
