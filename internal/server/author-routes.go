package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) getAuthors(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.GetAuthors())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func (s *Server) createAuthor(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.CreateAuthor(r))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func (s *Server) getAuthor(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.GetAuthor(r))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func (s *Server) updateAuthor(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.UpdateAuthor(r))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func (s *Server) deleteAuthor(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.DeleteAuthor(r))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}
