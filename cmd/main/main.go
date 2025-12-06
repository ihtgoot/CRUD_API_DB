package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ihtgoot/CRUD_API_DB/pkg/config"
	"github.com/ihtgoot/CRUD_API_DB/pkg/models"
	"github.com/ihtgoot/CRUD_API_DB/pkg/routs"
	"github.com/rs/cors"
)

func main() {
	db := config.Connect()
	models.Init(db)
	r := mux.NewRouter()
	routs.RegisterBookStoreRouts(r)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", c.Handler(r)))
}
