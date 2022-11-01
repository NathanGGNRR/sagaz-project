package models

//UserModule object for REST(CRUD)
type UserModule struct {
	UserId   int `json:"user_id"`
	ModuleId int `json:"module_id"`
}
