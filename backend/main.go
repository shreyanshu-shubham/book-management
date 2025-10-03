package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shreyanshu-shubham/book-management/backend/routes"
)

func main() {
	fmt.Println("Started....")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Authorization"})
	originOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})


	r := mux.NewRouter()
	routes.RegisterBookManagementRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9999", handlers.CORS(headersOk, originOk,methodsOk)(r)))
}
