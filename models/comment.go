package models

import (
	"time"
)

type Comment struct {
	ID       int        `gorm:"primary_key;auto_increment"`
	NewsID   int        `gorm:"column:news_id;type:int(11) not null"`
	Message  string     `gorm:"column:message;type:text;not null"`
	UserID   int        `gorm:"column:user_id;type:int(11) not null"`
	DatePost *time.Time `gorm:"column:date_post;type:datetime not null"`
}
