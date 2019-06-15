package models

type NewsTag struct {
	NewsID int `gorm:"column:news_id;primary_key;"`
	TagID  int `gorm:"column:tag_id;primary_key;"`
}
