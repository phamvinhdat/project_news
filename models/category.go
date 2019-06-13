package models

type Category struct {
	ID               int    `gorm:"primary_key;auto_increment"`
	Name             string `gorm:"column:name;type:nvarchar(255);not null"`
	ParentCategoryID int    `gorm:"column:parent_category_id;type:int(11)"`
}
