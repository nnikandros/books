package server

import (
	"books/internal/database"
	"net/http"
	"serde"
	"strconv"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
)

type BooksRouter struct {
	db        database.Service
	templates *template.Template
}

func (b *BooksRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", b.RenderBooksPage)
	r.Get("/{id}", b.RenderDetailsPage)

	r.Route("/api", func(r chi.Router) {
		r.Get("/", b.ListBooks)         // GET /books/api
		r.Post("/", b.CreateBook)       // POST /books/api
		r.Get("/{id}", b.GetBook)       // GET /books/{id}
		r.Delete("/{id}", b.DeleteBook) // DELETE /books/{id}  not implemented

	})

	return r
}

func (b *BooksRouter) ListBooks(w http.ResponseWriter, r *http.Request) {
	l, err := b.db.Queries.GetAllBooks(r.Context())
	if err != nil {
		http.Error(w, "error executing the query", http.StatusBadRequest)
	}

	err = serde.EncodeJson(w, http.StatusOK, l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (b *BooksRouter) CreateBook(w http.ResponseWriter, r *http.Request) {

	book, err := serde.DecodeV2[database.BookModel](r.Body)
	if err != nil {
		http.Error(w, "bad book model", http.StatusBadRequest)
		return
	}

	p, err := book.ParseAndValidate()
	if err != nil {
		http.Error(w, "bad book model", http.StatusBadRequest)
		return
	}

	err = b.db.Queries.AddBook(r.Context(), p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("created"))

}

func (b *BooksRouter) GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	book, err := b.db.Queries.GetBookById(r.Context(), int64(id))
	if err != nil {
		http.Error(w, "error while executign the query", http.StatusInternalServerError)
		return
	}
	err = serde.EncodeJson(w, http.StatusOK, book)
	if err != nil {
		http.Error(w, "error while serializing", http.StatusInternalServerError)
		return
	}
}

// not implemented  yet
func (b *BooksRouter) DeleteBook(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "bad it", http.StatusBadRequest)
		return
	}

	if err = b.db.Queries.DeleteBookById(r.Context(), int64(id)); err != nil {
		http.Error(w, "error deleting the book", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Deleted book ID: " + strconv.Itoa(id)))
}

func (b *BooksRouter) RenderBooksPage(w http.ResponseWriter, r *http.Request) {
	books, err := b.db.Queries.GetAllBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := b.templates.ExecuteTemplate(w, "books.html", books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (b *BooksRouter) RenderDetailsPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	book, err := b.db.Queries.GetBookById(r.Context(), int64(id))
	if err != nil {
		http.Error(w, "error while executign the query", http.StatusInternalServerError)
		return
	}

	if err := b.templates.ExecuteTemplate(w, "book_detail.html", book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func formatTime(t time.Time) string {
	return t.Format(time.DateOnly)

}
