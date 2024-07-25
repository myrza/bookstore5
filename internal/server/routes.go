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

	r.HandleFunc("/", s.HelloWorldHandler)

	r.HandleFunc("/health", s.healthHandler)
	r.HandleFunc("/create_db", s.dbStructure)

	r.HandleFunc("/authors", s.getAuthors)
	r.HandleFunc("/create_author", s.createAuthor)
	r.HandleFunc("/update_author/{id}", s.updateAuthor)
	r.HandleFunc("/author/{id}", s.getAuthor)
	r.HandleFunc("/delete_author/{id}", s.deleteAuthor)

	r.HandleFunc("/books", s.getBooks)
	r.HandleFunc("/create_book", s.createBook)
	r.HandleFunc("/update_book/{id}", s.updateBook)
	r.HandleFunc("/book/{id}", s.getBook)
	r.HandleFunc("/delete_book/{id}", s.deleteBook)

	r.HandleFunc("/book_author/{id}", s.getBookAndAuthor)
	r.HandleFunc("/update_book_author/{id}", s.updateBookAndAuthor)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func (s *Server) dbStructure(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.DbStructure())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Check if the request is for CORS preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass down the request to the next middleware (or final handler)
		next.ServeHTTP(w, r)
	})

}

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set JSON Content-Type
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
