package server

import (
	"log/slog"
	"os"

	"github.com/go-chi/httplog/v3"
)

func getLogger() *slog.Logger {

	logFormat := httplog.SchemaECS.Concise(false)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: logFormat.ReplaceAttr,
	})).With(
		slog.String("app", "books"),
		slog.String("version", "v1.0.0-a1fa420"),
		slog.String("env", "production"),
	)

	return logger

}
