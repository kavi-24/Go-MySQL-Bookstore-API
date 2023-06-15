package routes

import (
	"Bookstore/pkg/controllers"
	"github.com/gorilla/mux"
	// "github.com/akhil/go-bookstore/pkg/controllers"
	// "github.com/AkhilSharma90/Golang-MySQL-CRUD-Bookstore-Management-API/tree/master/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}