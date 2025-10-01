package routes

import (
	"github.com/gorilla/mux"
	"github.com/shreyanshu-shubham/book-management/backend/controllers"
)

var RegisterBookManagementRoutes = func(router *mux.Router) {
	router.HandleFunc("/addBook",controllers.AddBook).Methods("POST")
	router.HandleFunc("/getBooks",controllers.GetBooks).Methods("GET")
}