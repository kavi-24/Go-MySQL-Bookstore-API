package main

// go get "github.com/jinzhu/gorm"
// go get "github.com/jinzhu/gorm/dialects/mysql"
// go get "github.com/gorilla/mux"

import (
	"Bookstore/pkg/routes"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

