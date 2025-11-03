package server

import (
	"books/internal/database"
	"fmt"
	"log"
	"net/http"
	"serde"

	"github.com/go-chi/chi/v5"
)

type BooksRouter struct {
	db database.Service
}

func (b *BooksRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", b.ListBooks)         // GET /books
	r.Post("/", b.CreateBook)       // POST /books
	r.Get("/{id}", b.GetBook)       // GET /books/{id}
	r.Delete("/{id}", b.DeleteBook) // DELETE /books/{id}

	return r
}

func (b *BooksRouter) ListBooks(w http.ResponseWriter, r *http.Request) {
	l, err := b.db.Queries.GetAllBooks(r.Context())
	if err != nil {
		http.Error(w, "someting bad happend", http.StatusBadRequest)
	}

	err = serde.EncodeJson(w, http.StatusOK, l)
	fmt.Println(err)
}

func (b *BooksRouter) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := database.AddBookParams{}
	err := b.db.Queries.AddBook(r.Context(), book)
	if err != nil {
		log.Printf("b.db.Queries.AddBook %w", err)
	}
	w.Write([]byte("Book created"))
}

func (b *BooksRouter) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Book ID: " + id))
}

func (b *BooksRouter) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Deleted book ID: " + id))
}
