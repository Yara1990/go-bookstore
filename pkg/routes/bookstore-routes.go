package routes

import (
	"github.com/gorilla/mux"
	"github.com/yara1990/go-bookstore/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/boo/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book.{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/bookId", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/bookId", controllers.DeleteBook).Methods("DELETE")
}
