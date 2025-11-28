package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "net/http/pprof"

	_ "github.com/joho/godotenv/autoload"

	"books/internal/database"
)

type Server struct {
	port      int
	db        database.Service
	templates *template.Template
}

func New() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	t := template.Must(template.New("book_templates").Funcs(template.FuncMap{"formatTime": formatTime}).ParseGlob("templates/*"))

	NewServer := &Server{
		port:      port,
		db:        database.NewService(),
		templates: t,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
