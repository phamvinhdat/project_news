package models

type Tag struct{
	ID          int        `gorm:"primary_key;auto_increment"`
	Name        string     `gorm:"column:name;type:nvarchar(50);not null"`
}