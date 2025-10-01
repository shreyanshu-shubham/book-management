package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	db_utils "github.com/shreyanshu-shubham/book-management/backend/db-utils"
	"github.com/shreyanshu-shubham/book-management/backend/models"
	"github.com/shreyanshu-shubham/book-management/backend/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	include_trash, _ := strconv.ParseBool(r.URL.Query().Get("include_trash"))
	only_trash, _ := strconv.ParseBool(r.URL.Query().Get("only_trash"))

	var data []models.Book
	if include_trash {
		data = db_utils.GetAllBooks()
	} else if only_trash {
		data = db_utils.GetOnlyTrashedBooks()
	} else {
		data = db_utils.GetActiveBooks()
	}

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
func UpdateBook(w http.ResponseWriter, r *http.Request) {}
func TrashBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["isbn"]
	ISBN, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	db_utils.TrashBook(ISBN)
}
func RestoreBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["isbn"]
	ISBN, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	db_utils.RestoreBook(ISBN)
}
func ClearTrash(w http.ResponseWriter, r *http.Request)   {}
func RestoreTrash(w http.ResponseWriter, r *http.Request) {}
func DeleteBook(w http.ResponseWriter, r *http.Request)   {}
