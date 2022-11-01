package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

func main() {

	app := App{}
	app.initialiseRoutes()
	app.run()

	

}