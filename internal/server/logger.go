package server

import (
	"log/slog"
	"net/http"
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

var loggerOptions = &httplog.Options{

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
}
