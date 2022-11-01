package models

import "github.com/jinzhu/gorm"

//Module object for REST(CRUD)
type Module struct {
	gorm.Model
	Name          string      `json:"name" gorm:"column:name;not null;type:varchar(45)"`
	UserCreatorID int         `json:"user_creator_id"`
	UserCreator   User        `gorm:"foreignKey:UserCreatorID;references:ID"`
	Resources     []*Resource `gorm:"many2many:resource_modules;"`
	Participants  []*User     `gorm:"many2many:module_participants;"`
}
