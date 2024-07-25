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

func (s *service) CreateAuthor(r *http.Request) interface{} {
	//m := types.ReturnMessage{}

	var a types.Author
	json.NewDecoder(r.Body).Decode(&a)

	err := s.db.QueryRow("INSERT INTO authors (name, surname, biography, birthday) VALUES ($1, $2, $3, $4) RETURNING id", a.Name, a.Surname, a.Biography, a.Birthday).Scan(&a.ID)
	if err != nil {
		return err
	}

	//m.Message = "Автор успешно создан. Новый идентификатор"
	return a
}
func (s *service) GetAuthor(r *http.Request) interface{} {
	//m := types.ReturnMessage{}
	vars := mux.Vars(r)
	id := vars["id"]

	var a types.Author
	err := s.db.QueryRow("SELECT * FROM authors WHERE id = $1", id).Scan(&a.ID, &a.Name, &a.Surname, &a.Biography, &a.Birthday)
	if err != nil {
		return err
	}

	return a
}
func (s *service) UpdateAuthor(r *http.Request) interface{} {

	var a types.Author
	json.NewDecoder(r.Body).Decode(&a)

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := s.db.Exec("UPDATE authors SET name = $1, surname = $2,biography=$3, birthday = $4  WHERE id = $5", a.Name, a.Surname, a.Biography, a.Birthday, id)
	if err != nil {
		return err
	}

	return a
}
func (s *service) GetAuthors() interface{} {
	rows, err := s.db.Query("SELECT * FROM authors")
	if err != nil {
		return err
	}
	defer rows.Close()

	authors := []types.Author{}
	for rows.Next() {
		var a types.Author
		if err := rows.Scan(&a.ID, &a.Name, &a.Surname, &a.Biography, &a.Birthday); err != nil {

			return err
		}
		authors = append(authors, a)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return authors
}
func (s *service) DeleteAuthor(r *http.Request) interface{} {
	m := types.ReturnMessage{}

	var a types.Author
	json.NewDecoder(r.Body).Decode(&a)

	vars := mux.Vars(r)
	id := vars["id"]

	err := s.db.QueryRow("SELECT * FROM authors WHERE id = $1", id).Scan(&a.ID, &a.Name, &a.Surname, &a.Biography, &a.Birthday)
	if err != nil {
		m.Message = "Автор не обнаружен"
		return m
	} else {
		_, err := s.db.Exec("DELETE FROM authors WHERE id = $1", id)
		if err != nil {
			return err
		}

	}

	m.Message = "Автор удален"
	return m
}
