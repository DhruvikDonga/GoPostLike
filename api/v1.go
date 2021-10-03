package api

import (
	UserController "github.com/DhruvikDonga/goshopcart/controllers"

	"github.com/gorilla/mux"
)

func V1(r *mux.Router) {

	r.HandleFunc("/users", UserController.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", UserController.GetUser).Methods("GET")
	r.HandleFunc("/users", UserController.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UserController.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", UserController.DeleteUsers).Methods("DELETE")

	//logs

}
