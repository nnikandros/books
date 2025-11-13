package paths

import (
	"log"
	"path/filepath"
)

func appDir() string {
	b, err := filepath.Rel(".", "../../../books")
	if err != nil {
		log.Fatalf("failed to get relative path: %v", err)
	}
	abs, err := filepath.Abs(b)
	if err != nil {
		log.Fatalf("failed to get absolute path: %v", err)
	}
	return abs
}

var booksDir = appDir()

func SqliteTestFile() string {
	return filepath.Join(booksDir, "test.db")

}

func Favicon() string {
	return filepath.Join(booksDir, "favicon.ico")

}

func SqliteProdFile() string {
	return filepath.Join(booksDir, "prod.db")
}

func BooksJsonFile() string {
	return filepath.Join(booksDir, "books.json")
}
