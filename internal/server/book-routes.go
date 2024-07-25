package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) getBooks(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.GetBooks())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
func (s *Server) createBook(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.CreateBook(r))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
func (s *Server) getBook(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.GetBook(r))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
func (s *Server) updateBook(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.UpdateBook(r))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
func (s *Server) deleteBook(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.DeleteBook(r))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
