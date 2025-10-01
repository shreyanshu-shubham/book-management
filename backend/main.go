package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyanshu-shubham/book-management/backend/routes"
)

func main() {
	fmt.Println("Started....")
	r := mux.NewRouter()
	routes.RegisterBookManagementRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9999", r))
}
