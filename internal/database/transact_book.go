package database

import (
	"context"
	"encoding/json"
	"net/http"

	"bookstore5/internal/types"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func (s *service) GetBookAndAuthor(r *http.Request) interface{} {
	m := types.ReturnMessage{}
	vars := mux.Vars(r)
	id := vars["id"]

	var b types.AuthorAndBook

	err := s.db.QueryRow("SELECT name, surname, biography, birthday, title, authorid, isbn, year FROM books, authors  WHERE books.authorid = authors.id and  books.id = $1", id).Scan(&b.Name, &b.Surname, &b.Biography, &b.Birthday, &b.Title, &b.AuthorID, &b.ISBN, &b.Year)
	if err != nil {
		m.Message = "Ошибка при попытке получить книгу по указанному идентификатору"
		return m
	}

	return b
}

func (s *service) UpdateBookAndAuthor(r *http.Request) interface{} {
	m := types.ReturnMessage{}

	var ab types.AuthorAndBook

	json.NewDecoder(r.Body).Decode(&ab)

	vars := mux.Vars(r)
	book_id := vars["id"]

	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, "UPDATE books SET title = $1, authorid = $2, isbn=$3, year = $4  WHERE id = $5", ab.Title, ab.AuthorID, ab.ISBN, ab.Year, book_id)

	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = s.db.ExecContext(ctx, "UPDATE authors SET name = $1, surname = $2,biography=$3, birthday = $4  WHERE id = $5", ab.Name, ab.Surname, ab.Biography, ab.Birthday, ab.AuthorID)
	if err != nil {
		tx.Rollback()
		return err

	}
	err = tx.Commit()
	if err != nil {
		return err
	} else {
		m.Message = "Транзакция выполнена"
		return m
	}

}
