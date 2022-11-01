package database

import (
	"log"
	"sagaz-api/models"

	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

//Migrate create/updates database table
func Migrate() {
	Connector.Debug().AutoMigrate(&models.ResourceType{}, &models.Resource{})
	Connector.Debug().AutoMigrate(&models.UserType{}, &models.User{})
	Connector.Debug().AutoMigrate(&models.Module{})
	Connector.Debug().Model(&models.User{}).AddForeignKey("user_type_id", "users(user_type_id)", "CASCADE", "CASCADE")
	log.Println("Tables migrated")
}

func InitDB() {
	config :=
		Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "azerty",
			DB:         "sagaz",
		}

	connectionString := GetConnectionString(config)
	err := Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	Migrate()
}
