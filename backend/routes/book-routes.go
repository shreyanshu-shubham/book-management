package routes

import (
	"github.com/gorilla/mux"
	"github.com/shreyanshu-shubham/book-management/backend/controllers"
)

var RegisterBookManagementRoutes = func(router *mux.Router) {
	router.HandleFunc("/addBook", controllers.AddBook).Methods("POST")
	router.HandleFunc("/getBooks", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/getBook/{isbn}", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/trashBook/{isbn}", controllers.TrashBook).Methods("GET")
	router.HandleFunc("/restoreBook/{isbn}", controllers.RestoreBook).Methods("GET")
	router.HandleFunc("/deleteBook/{isbn}", controllers.DeleteBook).Methods("GET")


	router.HandleFunc("/restoreTrash", controllers.RestoreTrash).Methods("GET")
	router.HandleFunc("/emptyTrash", controllers.EmptyTrash).Methods("GET")
}
