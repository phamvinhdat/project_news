package models

type User struct {
	ID          int        `gorm:"primary_key;auto_increment"`
	Username    string     `form:"username" binding:"required" gorm:"column:username;type:varchar(50);not null"`
	Password    string     `form:"password" binding:"required" gorm:"column:password;type:varchar(255);not null;unique"`
	RoleID      int        `gorm:"column:role_id;type:varchar(255);not null"`
	Name        string     `form:"name" binding:"required" gorm:"column:name;type:nvarchar(50);not null"`
	PhoneNumber string     `gorm:"column:phone_number;type:varchar(12);not null"`
	Sex         bool       `form:"sex" binding:"required" gorm:"column:sex;type:bool;not null"`
	Email       string     `form:"email" binding:"required" gorm:"column:email;type:varchar(60);unique"`
}
