package server

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "net/http/pprof"

	_ "github.com/joho/godotenv/autoload"

	"books/internal/database"
)

type Server struct {
	host      string
	port      int
	db        database.Service
	templates *template.Template
	log       *slog.Logger
}

func New() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")
	t := template.Must(template.New("book_templates").Funcs(template.FuncMap{"formatTime": formatTime}).ParseGlob("templates/*"))

	NewServer := &Server{
		host:      host,
		port:      port,
		db:        database.NewService(),
		templates: t,
		log:       getLogger(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", NewServer.host, NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
