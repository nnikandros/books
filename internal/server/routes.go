package server

import (
	"books/internal/paths"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v3"
)

func (s *Server) RegisterRoutes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	// Request logger
	r.Use(httplog.RequestLogger(s.logger, &httplog.Options{

		Level: slog.LevelInfo,

		// Set log output to Elastic Common Schema (ECS) format.
		Schema: httplog.SchemaECS,

		// RecoverPanics recovers from panics occurring in the underlying HTTP handlers
		// and middlewares. It returns HTTP 500 unless response status was already set.
		//
		// NOTE: Panics are logged as errors automatically, regardless of this setting.
		RecoverPanics: true,

		// Optionally, filter out some request logs.
		Skip: func(req *http.Request, respStatus int) bool {

			faviconPath := req.URL.Path == "/favicon.ico"
			return respStatus == 404 || respStatus == 405 || faviconPath
		},

		// Optionally, log selected request/response headers explicitly.
		LogRequestHeaders:  []string{"Origin"},
		LogResponseHeaders: []string{},

		// Optionally, enable logging of request/response body based on custom conditions.
		// Useful for debugging payload issues in development.
		// LogRequestBody:  func(req *http.Request) bool { return false },
		// LogResponseBody: func(req *http.Request) bool { return true },
	}))

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

	r.Mount("/debug", middleware.Profiler())

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, paths.Favicon())
	})

	booksAPI := BooksAPIRouter{db: s.db}
	r.Mount("/api", booksAPI.Routes())

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	// panic("fuck")
	http.Error(w, "fuck off", http.StatusInternalServerError)
	return
	// jsonResp, _ := json.Marshal(s.db.Health())
	// w.Header().Set("content-type", "application/json")
	// _, _ = w.Write(jsonResp)
}

func (s *Server) RenderBooksPage(w http.ResponseWriter, r *http.Request) {
	books, err := s.db.Queries.GetAllBooksSortedByDate(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := s.templates.ExecuteTemplate(w, "books.html", books); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}

func (s *Server) RenderDetailsPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	book, err := s.db.Queries.GetBookById(r.Context(), int64(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := s.templates.ExecuteTemplate(w, "book_detail.html", book); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
