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

//GetAllResourceType get all resource_type data
func GetAllResourceType(w http.ResponseWriter, r *http.Request) {
	var resource_types []models.ResourceType
	database.Connector.Find(&resource_types)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resource_types)
}

//GetResourceTypeByID returns resource_type with specific ID
func GetResourceTypeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var resource_type models.ResourceType
	database.Connector.First(&resource_type, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource_type)
}

//CreateResourceType creates resource_type
func CreateResourceType(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var resource_type models.ResourceType
	json.Unmarshal(requestBody, &resource_type)

	database.Connector.Create(resource_type)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource_type)
}

//UpdateResourceTypeByID updates resource_type with respective ID
func UpdateResourceTypeByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var resource_type models.ResourceType
	json.Unmarshal(requestBody, &resource_type)
	database.Connector.Save(&resource_type)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resource_type)
}

//DeleteResourceTypeByID delete's resource_type with specific ID
func DeleteResourceTypeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var resource_type models.ResourceType
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&resource_type)
	w.WriteHeader(http.StatusNoContent)
}
