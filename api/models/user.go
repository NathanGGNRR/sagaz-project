package models

import "github.com/jinzhu/gorm"

//User object for REST(CRUD)
type User struct {
	gorm.Model
	Login      string `json:"login" gorm:"column:login;type:varchar(45);not null"`
	FirstName  string `json:"first_name" gorm:"first_name:login;type:varchar(45);not null"`
	LastName   string `json:"last_name" gorm:"column:last_name;type:varchar(45);not null"`
	UserTypeID int
}
