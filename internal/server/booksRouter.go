package server

import (
	"books/internal/database"
	"net/http"

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
	w.Write([]byte("List of books"))
}

func (b *BooksRouter) CreateBook(w http.ResponseWriter, r *http.Request) {
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
