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
		log.Fatal("Failed to connect to DB")
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Ping failed to DB")
		log.Fatal(err)
	}
	return db
}

func AddBook(book models.Book) bool {
	db := getConnection()
	defer db.Close()

	log.Printf("inserting book: %+v\n",book)

	_, st := GetBookByISBN(book.ISBN)
	if st {
		log.Fatal("Duplice ISBN error")
		return false
	}

	var isbn int64
	query := `insert into books (isbn, title, author, is_deleted) values ($1,$2,$3,$4) returning isbn`
	err := db.QueryRow(query, book.ISBN, book.Title, book.Author, book.IsDeleted).Scan(&isbn)
	if err != nil {
		log.Fatal("Failed to insert")
		log.Fatal(err)
		return false
	}

	log.Printf("inserted successfully")
	return true
}

func TrashBook(isbn int64) bool {
	db := getConnection()
	defer db.Close()

	log.Printf("moving %d to trash",isbn)

	_,st := GetBookByISBN(isbn)
	if !st {
		log.Fatal("book with isbn(",isbn,") does not exist")
		return false
	}

	_, err := db.Exec(
		"update books set is_deleted=true where isbn=$1",
		isbn,
	)
	if err != nil {
		log.Fatal("failed to trash book(isbn=",isbn,")")
		return false
	}

	return true
}

func RestoreBook(isbn int64) bool {
	db := getConnection()
	defer db.Close()

	log.Printf("restoring book(isbn=%d)",isbn)

	_,st := GetBookByISBN(isbn)
	if !st {
		log.Fatalf("book (isbn=%d) does not exist",isbn)
		return false
	}

	_, err := db.Exec(
		"update books set is_deleted=false where isbn=$1",
		isbn,
	)

	if err != nil {
		log.Fatalf("failed to restore book(isbn=%d)",isbn)
		return false
	}

	return true
}

func DeleteBook(isbn int64) bool {
	db := getConnection()
	defer db.Close()

	_, err := db.Exec(
		"delete from books where isbn=$1 and is_deleted=true",
		isbn,
	)

	if err != nil {
		log.Fatalf("failed to delete book(isbn=%d)",isbn)
		return false
	}
	return true
}

func GetAllBooks() []models.Book {
	db := getConnection()
	defer db.Close()

	data := []models.Book{}
	rows, err := db.Query("select isbn,title,author,is_deleted from books")
	if err != nil {
		log.Fatalf("failed to get all book from DB")
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
			log.Fatalf("failed to scan row")
			log.Fatal(err)
		}
		data = append(data, models.Book{Title: title, ISBN: isbn, Author: author, IsDeleted: is_deleted})
	}

	log.Printf("got %d books from DB",len(data))
	return data
}

func GetOnlyTrashedBooks() []models.Book {
	db := getConnection() 
	defer db.Close()

	data := []models.Book{}
	rows, err := db.Query("select isbn,title,author,is_deleted from books where is_deleted=true")
	if err != nil {
		log.Fatalf("failed to get trashed book list from DB")
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
			log.Fatal("failed to scan row")
			log.Fatal(err)
		}
		data = append(data, models.Book{Title: title, ISBN: isbn, Author: author, IsDeleted: is_deleted})
	}

	log.Printf("fetched total of %d trashed books",len(data))
	return data
}

func GetActiveBooks() []models.Book {
	db := getConnection()
	defer db.Close()

	data := []models.Book{}
	rows, err := db.Query("select isbn,title,author,is_deleted from books where is_deleted=false")
	if err != nil {
		log.Fatal("failed to get books from DB")
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
			log.Fatal("failed to scan row")
			log.Fatal(err)
		}
		data = append(data, models.Book{Title: title, ISBN: isbn, Author: author, IsDeleted: is_deleted})
	}

	log.Printf("fetched total of %d active books",len(data))
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
			log.Fatalf("no book with isbn=%d",isbn)
			return models.Book{}, false
		}
		log.Fatalf("failed to fetch from DB")
		log.Fatal(err)
	}

	data := models.Book{Title: title, ISBN: isbn, Author: author, IsDeleted: is_deleted}
	log.Printf("got book from DB: %+v",data)
	return data, true
}

func RestoreTrash() bool {
	db := getConnection()
	defer db.Close()

	_, err := db.Exec(
		"update books set is_deleted=false where is_deleted=true",
	)

	if err != nil {
		log.Fatalf("failed to restore books from the trash in")
		log.Fatal(err)
	}
	return true
}

func EmptyTrash() bool {
	db := getConnection()
	defer db.Close()

	_, err := db.Exec(
		"delete from books where is_deleted=true",
	)
	if err != nil {
		log.Fatalf("failed to delete books from the trash in DB")
		log.Fatal(err)
	}
	return true
}

func UpdateBook(book models.Book) (bool) {
	db := getConnection()
	defer db.Close()

	

	_, st := GetBookByISBN(book.ISBN)
	if !st {
		return false
	}

	query := `update books set title=$1, author=$2 where isbn=$3`
	_, err := db.Exec(query, book.Title, book.Author, book.ISBN)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}