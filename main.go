package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	dbs "github.com/DhruvikDonga/goshopcart/DBs"
	"github.com/DhruvikDonga/goshopcart/api"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	apirouters := r.PathPrefix("/api").Subrouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./storage/postimage"))) //image routes
	dbs.IntialMigration()
	api.V1(apirouters)
	corshandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
			"POST",
		},

		AllowedHeaders: []string{
			"Token", //or you can your header key values which you are using in your application
			"X-Requested-With", "Content-Type",
		},
	}).Handler(r)
	srv := &http.Server{
		Addr:         "localhost:9000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      corshandler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)

}
