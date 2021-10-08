package api

import (
	Homecontrollers "github.com/DhruvikDonga/goshopcart/controllers/HomeControllers"
	"github.com/DhruvikDonga/goshopcart/controllers/UserControllers"

	usermiddleware "github.com/DhruvikDonga/goshopcart/middlewares"

	"github.com/gorilla/mux"
)

//V1 routers deal with the api points under version 1 api group
func V1(r *mux.Router) {
	//Home routers
	r.HandleFunc("/welcome", Homecontrollers.Home).Methods("GET")

	//Users signin and signup
	r.HandleFunc("/signin", UserControllers.SignIn).Methods("POST")
	r.HandleFunc("/signup", UserControllers.SignUp).Methods("POST")

	r.HandleFunc("/users", usermiddleware.UserAuthorization(UserControllers.GetUsers)).Methods("GET")
	r.HandleFunc("/user/{id}", UserControllers.GetUser).Methods("GET")
	r.HandleFunc("/users", UserControllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UserControllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", UserControllers.DeleteUsers).Methods("DELETE")

	//logs

}
