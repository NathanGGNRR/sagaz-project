package models

//Resource object for REST(CRUD)
type Resource struct {
	ID            int             `gorm:"primaryKey"`
	Name          string          `json:"name" gorm:"column:name;not null;type:varchar(45)"`
	ResourceTypes []*ResourceType `gorm:"many2many:resource_types;"`
}
