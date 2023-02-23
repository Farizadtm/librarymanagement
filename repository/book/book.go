package bookRepo

import (
	"database/sql"
	"fmt"
	"librarysystem/structs"
)

func GetAllBooks(db *sql.DB) (result []structs.Book, err error) {
	sql := "SELECT * FROM book"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Release_year, &book.Category, &book.IsAvailable)
		if err != nil {
			panic(err)
		}

		result = append(result, book)
	}

	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := "INSERT INTO book (title, author, release_year, category, isavailable) VALUES ($1, $2, $3, $4, $5)"
	errs := db.QueryRow(sql, book.Title, book.Author, book.Release_year, book.Category, book.IsAvailable)
	return errs.Err()
}

func UpdateStatusBook(db *sql.DB, book structs.Book) (err error) {
	fmt.Println(book.IsAvailable)
	sql := "UPDATE book SET isavailable = $2 WHERE id = $1"
	errs := db.QueryRow(sql, book.ID, book.IsAvailable)
	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := "UPDATE book SET title = $2, author = $3, release_year = $4, category = $5 WHERE id = $1"
	errs := db.QueryRow(sql, book.ID, book.Title, book.Author, book.Release_year, book.Category)
	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM book where id = $1"
	errs := db.QueryRow(sql, book.ID)
	return errs.Err()
}

func GetBookByID(db *sql.DB, book structs.Book) (result structs.Book, err error) {
	sql := "SELECT * FROM book WHERE id = $1 AND isavailable = $2 LIMIT 1"

	rows, err := db.Query(sql, book.ID, book.IsAvailable)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Release_year, &book.Category, &book.IsAvailable)
		if err != nil {
			panic(err)
		}

		result = book
	}

	return
}
