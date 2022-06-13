package routes

import (
	controllers "github.com/XCross01/go-bookstore/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookId).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookId).Methods("DELETE")
}
