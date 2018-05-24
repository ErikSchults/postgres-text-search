package models

import (
	"fmt"
	"strings"

	"github.com/erikschults/go-postgre-text-search/app/database"
)

type Book struct {
	ID     uint
	Author string
	Notes  string
	Title  string
}

func AllBooks() ([]*Book, error) {
	return FetchBooks("SELECT id, author, notes, title FROM books")
}

func FullTextSearch(search string) ([]*Book, error) {
	query := fmt.Sprintf(`SELECT id, author, notes, title FROM books 
		WHERE document @@ to_tsquery($1)`)

	return FetchBooks(query, strings.Join(strings.Split(search, " "), " & "))
}

func FetchBooks(query string, args ...interface{}) ([]*Book, error) {
	rows, err := database.DBCon.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]*Book, 0)
	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.ID, &book.Author, &book.Notes, &book.Title)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
