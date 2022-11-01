package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sagaz-api/database"
	"sagaz-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllModule get all module data
func GetAllModule(w http.ResponseWriter, r *http.Request) {
	var modules []models.Module
	database.Connector.Find(&modules)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(modules)
}

//GetModuleByID returns module with specific ID
func GetModuleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var module models.Module
	database.Connector.First(&module, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(module)
}

//CreateModule creates module
func CreateModule(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var module models.Module
	json.Unmarshal(requestBody, &module)

	database.Connector.Create(module)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(module)
}

//UpdateModuleByID updates module with respective ID
func UpdateModuleByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var module models.Module
	json.Unmarshal(requestBody, &module)
	database.Connector.Save(&module)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(module)
}

//DeleteModuleByID delete's module with specific ID
func DeleteModuleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var module models.Module
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&module)
	w.WriteHeader(http.StatusNoContent)
}
