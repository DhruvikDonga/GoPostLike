package main

import (
	"log"
	"net/http"

	dbs "github.com/DhruvikDonga/goshopcart/DBs"
	"github.com/DhruvikDonga/goshopcart/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	dbs.IntialMigration()
	api.V1(r)

	log.Fatal(http.ListenAndServe(":9000", r))

}
