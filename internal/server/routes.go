package server

import (
	"books/internal/paths"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", s.RenderBooksPage)
	r.Get("/{id}", s.RenderDetailsPage)

	r.Get("/health", s.healthHandler)

	// r.Mount("/debug", middleware.Profiler())

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, paths.Favicon())
	})

	booksAPI := BooksAPIRouter{db: s.db}
	r.Mount("/api", booksAPI.Routes())

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(jsonResp)
}

func (s *Server) RenderBooksPage(w http.ResponseWriter, r *http.Request) {
	books, err := s.db.Queries.GetAllBooksSortedByDate(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.templates.ExecuteTemplate(w, "books.html", books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (s *Server) RenderDetailsPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	book, err := s.db.Queries.GetBookById(r.Context(), int64(id))
	if err != nil {
		http.Error(w, "error while executign the query", http.StatusInternalServerError)
		return
	}

	if err := s.templates.ExecuteTemplate(w, "book_detail.html", book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
