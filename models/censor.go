package models

import "time"

type Censor struct {
	NewsID     int        `gorm:"column:news_id;primary_key;"`
	UserID     int        `gorm:"column:user_id;primary_key;"`
	IsPublic   bool       `gorm:"colums:is_public;type:bool;default false"`
	DateCensor *time.Time `gorm:"colums:date_censor;type:timestamp;not null"`
	DatePublic *time.Time `gorm:"colums:date_public;type:timestamp;not null"`
}
