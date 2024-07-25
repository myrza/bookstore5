package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) getBookAndAuthor(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.GetBookAndAuthor(r))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) updateBookAndAuthor(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.UpdateBookAndAuthor(r))

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
