package routes

import (
	"github.com/gorilla/mux"
	"github.com/iamsushank/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/topTen", controllers.GetTopTenBook).Methods("GET")
	router.HandleFunc("/book/{author}/{pub}", controllers.GetBookByAuthorAndName).Methods("GET")
	router.HandleFunc("/book/struct", controllers.GetBookUsingStruct).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
