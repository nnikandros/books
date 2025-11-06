package server

import (
	"books/internal/database"
	"fmt"
	"net/http"
	"serde"
	"strconv"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
)

type BooksRouter struct {
	db database.Service
}

func (b *BooksRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", b.RenderBooksPage)
	r.Get("/{id}", b.RenderDetalsPage)

	r.Route("/api", func(r chi.Router) {
		r.Get("/", b.ListBooks)         // GET /books
		r.Post("/", b.CreateBook)       // POST /books
		r.Get("/{id}", b.GetBook)       // GET /books/{id}
		r.Delete("/{id}", b.DeleteBook) // DELETE /books/{id}

	})

	return r
}

func (b *BooksRouter) ListBooks(w http.ResponseWriter, r *http.Request) {
	l, err := b.db.Queries.GetAllBooks(r.Context())
	if err != nil {
		http.Error(w, "someting bad happend", http.StatusBadRequest)
	}

	// sliceBooks := make([]database.BookModel, 0, len(l))
	// for _, i := range l {
	// 	sliceBooks = append(sliceBooks, database.FromBook(i))
	// }

	err = serde.EncodeJson(w, http.StatusOK, l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (b *BooksRouter) CreateBook(w http.ResponseWriter, r *http.Request) {

	book, err := serde.DecodeV2[database.BookModel](r.Body)
	if err != nil {
		fmt.Printf("serde.DecodeV %v", err)
		http.Error(w, "bad book model", http.StatusBadRequest)
		return
	}

	p, err := book.ParseBookModel()
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
	id := chi.URLParam(r, "id")
	id2, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	book, err := b.db.Queries.GetBookById(r.Context(), int64(id2))
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
	id := chi.URLParam(r, "id")
	w.Write([]byte("Deleted book ID: " + id))
}

func (b *BooksRouter) RenderBooksPage(w http.ResponseWriter, r *http.Request) {
	books, err := b.db.Queries.GetAllBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/books.html"))
	err = tmpl.Execute(w, map[string]any{
		"Title": "Books List",
		"Books": books,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (b *BooksRouter) RenderDetalsPage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id2, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	book, err := b.db.Queries.GetBookById(r.Context(), int64(id2))
	if err != nil {
		http.Error(w, "error while executign the query", http.StatusInternalServerError)
		return
	}

	// tmpl := template.Must(template.ParseFiles("templates/book_detail.html")).Funcs(template.FuncMap{"formatTime": formatTime})
	tmpl := template.Must(template.New("book_detail.html").Funcs(template.FuncMap{"formatTime": formatTime}).ParseFiles("templates/book_detail.html"))

	err = tmpl.Execute(w, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formatTime(t time.Time) string {
	return t.Format(time.DateOnly)

}
