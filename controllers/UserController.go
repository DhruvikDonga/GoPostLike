package controllers

import (
	"encoding/json"
	"net/http"

	dbs "github.com/DhruvikDonga/goshopcart/DBs"
	"github.com/DhruvikDonga/goshopcart/models"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	var users []models.User
	dbs.DB.Find(&users)
	json.NewEncoder(w).Encode(users)

}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	dbs.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	dbs.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	dbs.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	dbs.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}
func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	dbs.DB.Delete(&user, params["id"])

	//permanent delte the record
	//dbs.DB.Unscoped().Delete(&user, params["id"])

	json.NewEncoder(w).Encode("The user is deleted successfully!")
}
