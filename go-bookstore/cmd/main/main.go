package main

import (
	"github.com/gorilla/mux"
	"github.com/iamsushank/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8710", r))
}
