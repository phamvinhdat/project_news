package models

import (
	"time"
)

type News struct {
	ID         int        `gorm:"primary_key;auto_increment"`
	Title      string     `form:"title" binding:"required" gorm:"column:title;type:nvarchar(255);not null"`
	Avatar     string     `gorm:"column:avatar;type:nvarchar(255); not null"`
	Summary    string     `form:"summary" binding:"required" gorm:"column:summary;type:text;not null"`
	Content    string     `form:"content" binding:"required" gorm:"column:content;type:longtext;not null"`
	UserID     int        `gorm:"column:user_id;type:int(11) not null"`
	DatePost   *time.Time `time_format:"2006-01-02" time_utc:"1" gorm:"column:date_post;type:datetime not null"`
	CategoryID int        `form:"category" binding:"required" gorm:"column:category_id;type:int(11) not null"`
	Views      int        `gorm:"colums:views;type:int(11);default 0"`
	IsPremium  bool       `from:"ispremium" gorm:"colums:is_premium;type:bool;default 0"`
}
