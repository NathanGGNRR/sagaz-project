package main

import (
	"log"
	"net/http"
	"sagaz-api/database"
	"sagaz-api/controllers"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) initialiseRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/resources", controllers.GetAllResource).Methods("GET")
	app.Router.HandleFunc("/resource/{id}", controllers.GetResourceByID).Methods("GET")
	app.Router.HandleFunc("/resource", controllers.CreateResource).Methods("POST")
	app.Router.HandleFunc("/resource_types", controllers.GetAllResourceType).Methods("GET")
	app.Router.HandleFunc("/resource_type", controllers.CreateResourceType).Methods("POST")
	app.Router.HandleFunc("/modules", controllers.GetAllModule).Methods("GET")
	app.Router.HandleFunc("/module", controllers.CreateModule).Methods("POST")
	app.Router.HandleFunc("/users", controllers.GetAllUser).Methods("GET")
	app.Router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	app.Router.HandleFunc("/user_types", controllers.GetAllUserType).Methods("GET")
	app.Router.HandleFunc("/user_type", controllers.CreateUserType).Methods("POST")
	app.Router.HandleFunc("/vm", controllers.CreateVM).Methods("POST")
	app.Router.HandleFunc("/vm/{vm_name}", controllers.GetVM).Methods("GET")
	app.Router.HandleFunc("/jenkins", controllers.JenkinsBuild).Methods("POST")
	app.Router.HandleFunc("/vminfos", controllers.GetInfos).Methods("GET")
	http.Handle("/", app.Router)
}

func (app *App) run() {
	database.InitDB()
	log.Println("Starting the HTTP server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
