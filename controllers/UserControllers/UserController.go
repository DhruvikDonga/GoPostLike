package UserControllers

import (
	"encoding/json"
	"log"
	"net/http"

	dbs "github.com/DhruvikDonga/goshopcart/DBs"
	usermiddleware "github.com/DhruvikDonga/goshopcart/middlewares"
	"github.com/DhruvikDonga/goshopcart/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	var getuser models.User
	dbs.DB.Model(&models.User{}).Where("email=?", user.Email).First(&getuser) //check if emails clashed between users
	if getuser.Email != "" {
		json.NewEncoder(w).Encode("Problem with email already in use!")
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)

	if err != nil {
		log.Fatalln("error in password hash")
	}
	user.Role = "user"
	dbs.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content- Type", "application/json")
	var logindetails models.Authentication
	err := json.NewDecoder(r.Body).Decode(&logindetails)

	if err != nil {
		json.NewEncoder(w).Encode("Error with decoding user")
		return
	}

	var getuser models.User

	dbs.DB.Model(&models.User{}).Where("email=?", logindetails.Email).First(&getuser)
	if getuser.Email == "" {
		json.NewEncoder(w).Encode("Error with email address")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(getuser.Password), []byte(logindetails.Password))

	if err != nil {
		json.NewEncoder(w).Encode("Error with password not found")
		return
	}

	validToken, err := usermiddleware.JWTgenerate(getuser.Email, getuser.Role)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	var token models.Token
	token.Email = getuser.Email
	token.Role = getuser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(token)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not authorized."))
		return
	}
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
