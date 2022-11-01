package models

import "github.com/jinzhu/gorm"

//UserType object for REST(CRUD)
type UserType struct {
	gorm.Model
	Name  string `json:"name" gorm:"column:name;type:varchar(45);not null"`
	Users []User `gorm:"foreignKey:UserTypeID;"`
}
