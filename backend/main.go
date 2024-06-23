package main

import (
	"api/handlers"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	//"github.com/myrza/bookstore5/backend/handlers"
)

// main function
func main() {
	//connect to database
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", "postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// логируем в файл
	flog, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer flog.Close()

	log.SetOutput(flog)

	// create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	// create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS authors (id SERIAL PRIMARY KEY, name TEXT, surname TEXT, biography TEXT, birthday DATE)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books (id SERIAL PRIMARY KEY, title TEXT, authorid INTEGER, isbn TEXT, year INTEGER)")
	if err != nil {
		log.Fatal(err)
	}

	// create router
	router := mux.NewRouter()

	// операции по авторам книг
	router.HandleFunc("/api/go/authors", handlers.GetAuthors(db)).Methods("GET")
	router.HandleFunc("/api/go/authors", handlers.CreateAuthor(db)).Methods("POST")
	router.HandleFunc("/api/go/authors/{id}", handlers.GetAuthor(db)).Methods("GET")
	router.HandleFunc("/api/go/authors/{id}", handlers.UpdateAuthor(db)).Methods("PUT")
	router.HandleFunc("/api/go/authors/{id}", handlers.DeleteAuthor(db)).Methods("DELETE")

	// операции по книгам
	router.HandleFunc("/api/go/books", handlers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/api/go/books", handlers.CreateBook(db)).Methods("POST")
	router.HandleFunc("/api/go/books/{id}", handlers.GetBook(db)).Methods("GET")
	router.HandleFunc("/api/go/books/{id}", handlers.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/api/go/books/{id}", handlers.DeleteBook(db)).Methods("DELETE")

	// транзакция
	router.HandleFunc("/api/go/book-author/{id}", handlers.UpdateBookAndAuthor(db)).Methods("PUT")
	router.HandleFunc("/api/go/book-author/{id}", handlers.GetBookAndAuthor(db)).Methods("GET")

	// wrap the router with CORS and JSON content type middlewares
	enhancedRouter := handlers.EnableCORS(handlers.JsonContentTypeMiddleware(router))

	// start server
	log.Fatal(http.ListenAndServe(":8000", enhancedRouter))
}

/*

func enableCORS(next http.Handler) http.Handler {
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

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set JSON Content-Type
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

*/
