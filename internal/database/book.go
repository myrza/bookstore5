package database

import (
	"encoding/json"
	"net/http"

	"bookstore5/internal/types"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func (s *service) CreateBook(r *http.Request) interface{} {
	var b types.Book
	json.NewDecoder(r.Body).Decode(&b)

	err := s.db.QueryRow("INSERT INTO books (title, authorid, isbn, year) VALUES ($1, $2, $3, $4) RETURNING id", b.Title, b.AuthorID, b.ISBN, b.Year).Scan(&b.ID)
	if err != nil {
		return err
	}

	//m.Message = "Автор успешно создан. Новый идентификатор"
	return b
}
func (s *service) GetBook(r *http.Request) interface{} {
	//m := types.ReturnMessage{}
	vars := mux.Vars(r)
	id := vars["id"]

	var b types.Book
	err := s.db.QueryRow("SELECT * FROM books WHERE id = $1", id).Scan(&b.ID, &b.Title, &b.AuthorID, &b.ISBN, &b.Year)
	if err != nil {
		return err
	}

	return b
}
func (s *service) UpdateBook(r *http.Request) interface{} {

	var b types.Book
	json.NewDecoder(r.Body).Decode(&b)

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := s.db.Exec("UPDATE books SET title = $1, authorid = $2, isbn=$3, year = $4  WHERE id = $5", b.Title, b.AuthorID, b.ISBN, b.Year, id)

	if err != nil {
		return err
	}

	return b
}
func (s *service) GetBooks() interface{} {
	rows, err := s.db.Query("SELECT * FROM books")
	if err != nil {
		return err
	}
	defer rows.Close()

	books := []types.Book{}
	for rows.Next() {
		var b types.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.AuthorID, &b.ISBN, &b.Year); err != nil {

			return err
		}
		books = append(books, b)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return books
}
func (s *service) DeleteBook(r *http.Request) interface{} {
	m := types.ReturnMessage{}

	var b types.Book
	json.NewDecoder(r.Body).Decode(&b)

	vars := mux.Vars(r)
	id := vars["id"]

	err := s.db.QueryRow("SELECT * FROM books WHERE id = $1", id).Scan(&b.ID, &b.Title, &b.AuthorID, &b.ISBN, &b.Year)
	if err != nil {
		m.Message = "Книга не обнаружена"
		return m
	} else {
		_, err := s.db.Exec("DELETE FROM books WHERE id = $1", id)
		if err != nil {
			return err
		}

	}

	m.Message = "Книга удалена"
	return m
}
