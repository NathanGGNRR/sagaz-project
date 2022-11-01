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

//GetAllResource get all resource data
func GetAllResource(w http.ResponseWriter, r *http.Request) {
	var resources []models.Resource
	database.Connector.Find(&resources)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resources)
}

//GetResourceByID returns resource with specific ID
func GetResourceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var resource models.Resource
	database.Connector.First(&resource, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource)
}

//CreateResource creates resource
func CreateResource(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var resource models.Resource
	json.Unmarshal(requestBody, &resource)

	database.Connector.Create(resource)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

//UpdateResourceByID updates resource with respective ID
func UpdateResourceByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var resource models.Resource
	json.Unmarshal(requestBody, &resource)
	database.Connector.Save(&resource)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resource)
}

//DeleteResourceByID delete's resource with specific ID
func DeleteResourceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var resource models.Resource
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&resource)
	w.WriteHeader(http.StatusNoContent)
}
