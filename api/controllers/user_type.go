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

//GetAllUserType get all user_type data
func GetAllUserType(w http.ResponseWriter, r *http.Request) {
	var user_types []models.UserType
	database.Connector.Find(&user_types)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user_types)
}

//GetUserTypeByID returns user_type with specific ID
func GetUserTypeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var user_type models.UserType
	database.Connector.First(&user_type, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user_type)
}

//CreateUserType creates user_type
func CreateUserType(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user_type models.UserType
	json.Unmarshal(requestBody, &user_type)

	database.Connector.Create(user_type)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user_type)
}

//UpdateUserTypeByID updates user_type with respective ID
func UpdateUserTypeByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user_type models.UserType
	json.Unmarshal(requestBody, &user_type)
	database.Connector.Save(&user_type)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user_type)
}

//DeleteUserTypeByID delete's user_type with specific ID
func DeleteUserTypeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var user_type models.UserType
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&user_type)
	w.WriteHeader(http.StatusNoContent)
}
