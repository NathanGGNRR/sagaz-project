package models

import "github.com/jinzhu/gorm"

//ResourceType object for REST(CRUD)
type ResourceType struct {
	gorm.Model
	Name      string      `json:"name" gorm:"column:name;not null;type:varchar(45)"`
	Resources []*Resource `gorm:"many2many:resources_types;"`
}
