package database

import (
	"database/sql"

	"net/http"

	"fmt"
	"log"

	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	DbStructure() interface{}
	//Health() map[string]string

	//getAuthorsNewVer() interface{}

	GetAuthors() interface{}
	GetAuthor(r *http.Request) interface{}
	CreateAuthor(r *http.Request) interface{}
	UpdateAuthor(r *http.Request) interface{}
	DeleteAuthor(r *http.Request) interface{}

	GetBooks() interface{}
	GetBook(r *http.Request) interface{}
	CreateBook(r *http.Request) interface{}
	UpdateBook(r *http.Request) interface{}
	DeleteBook(r *http.Request) interface{}

	GetBookAndAuthor(r *http.Request) interface{}
	UpdateBookAndAuthor(r *http.Request) interface{}

	Close() error
}

type service struct {
	db *sql.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)

	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}

func (s *service) DbStructure() interface{} {
	// create table if not exists
	type message struct {
		Message string `json:"Message"`
	}
	m := message{}
	_, err := s.db.Exec("CREATE TABLE IF NOT EXISTS authors (id SERIAL PRIMARY KEY, name TEXT, surname TEXT, biography TEXT, birthday DATE)")
	if err != nil {
		return err
	}
	_, err = s.db.Exec("CREATE TABLE IF NOT EXISTS books (id SERIAL PRIMARY KEY, title TEXT, authorid INTEGER, isbn TEXT, year INTEGER)")
	if err != nil {
		return err
	}
	m.Message = "Структура базы данных создана"
	return m
}
