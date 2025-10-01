package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shreyanshu-shubham/book-management/backend/controllers/utils"
	db_utils "github.com/shreyanshu-shubham/book-management/backend/db-utils"
	"github.com/shreyanshu-shubham/book-management/backend/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	data := db_utils.GetAllBooks()
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBook(w http.ResponseWriter, r *http.Request) {}
func AddBook(w http.ResponseWriter, r *http.Request) {
	b := models.Book{}
	utils.ParseBody(r, &b)
	res := db_utils.AddBook(b)
	if !res {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func UpdateBook(w http.ResponseWriter, r *http.Request)   {}
func TrashBook(w http.ResponseWriter, r *http.Request)    {}
func RestoreBook(w http.ResponseWriter, r *http.Request)  {}
func ClearTrash(w http.ResponseWriter, r *http.Request)   {}
func RestoreTrash(w http.ResponseWriter, r *http.Request) {}
func DeleteBook(w http.ResponseWriter, r *http.Request)   {}
