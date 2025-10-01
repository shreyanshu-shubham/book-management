package db_utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/shreyanshu-shubham/book-management/backend/models"
)

func getConnection() *sql.DB {
	connectionString := "postgres://postgres:postgres@localhost:5432/bkmanage?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func AddBook(book models.Book) bool {
	db := getConnection()
	defer db.Close()

	_, st := GetBookByISBN(book.ISBN)
	if st {
		return false
	}

	var isbn int64
	query := `insert into books (isbn, title, author, is_deleted) values ($1,$2,$3,$4) returning isbn`
	err := db.QueryRow(query, book.ISBN, book.Title, book.Author, book.IsDeleted).Scan(&isbn)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func TrashBook(isbn int64) bool {
	db := getConnection()
	defer db.Close()

	_,st := GetBookByISBN(isbn)
	if !st {
		return false
	}

	_, err := db.Exec(
		"update books set is_deleted=true where isbn=$1",
		isbn,
	)
	return err == nil
}

func RestoreBook(isbn int64) bool {
	db := getConnection()
	defer db.Close()

	_,st := GetBookByISBN(isbn)
	if !st {
		return false
	}

	_, err := db.Exec(
		"update books set is_deleted=false where isbn=$1",
		isbn,
	)
	return err == nil
}

func DeleteBok() {}

func GetAllBooks() []models.Book {
	db := getConnection()
	defer db.Close()

	data := []models.Book{}
	rows, err := db.Query("select isbn,title,author,is_deleted from books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var isbn int64
	var title string
	var author string
	var is_deleted bool

	for rows.Next() {
		err := rows.Scan(&isbn, &title, &author, &is_deleted)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, models.Book{Title: title, ISBN: isbn, Author: author, IsDeleted: is_deleted})
	}

	return data
}

func GetBookByISBN(isbn int64) (models.Book, bool) {
	db := getConnection()
	defer db.Close()

	query := `select title, author, is_deleted from books where isbn=$1`
	var title string
	var author string
	var is_deleted bool
	err := db.QueryRow(query, isbn).Scan(&title, &author, &is_deleted)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, false
		}
		log.Fatal(err)
	}
	return models.Book{Title: title, ISBN: isbn, Author: author, IsDeleted: is_deleted}, true
}
