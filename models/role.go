package models

type Role struct{
	ID          int        `gorm:"primary_key;auto_increment"`
	Name        string     `gorm:"column:name;type:nvarchar(255);not null"`
}