package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func (s *Server) RegisterRoutes() http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/create_db", s.dbStructure).Methods("POST")

	r.HandleFunc("/authors", s.getAuthors).Methods("GET")
	r.HandleFunc("/authors", s.createAuthor).Methods("POST")
	r.HandleFunc("/authors/{id}", s.updateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", s.getAuthor).Methods("GET")
	r.HandleFunc("/authors/{id}", s.deleteAuthor).Methods("DELETE")

	r.HandleFunc("/books", s.getBooks).Methods("GET")
	r.HandleFunc("/books", s.createBook).Methods("POST")
	r.HandleFunc("/books/{id}", s.updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", s.getBook).Methods("GET")
	r.HandleFunc("/books/{id}", s.deleteBook).Methods("DELETE")

	r.HandleFunc("/book_author/{id}", s.getBookAndAuthor).Methods("GET")
	r.HandleFunc("/book_author/{id}", s.updateBookAndAuthor).Methods("PUT")

	return r

}

func (s *Server) dbStructure(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.DbStructure())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}
