package paths

import (
	"path/filepath"
	"runtime"
)

func rootDir() string {
	// Get file path of this file
	_, filename, _, _ := runtime.Caller(0)

	// filename = ".../yourproject/paths/paths.go"
	// root = parent of parent of that dir
	return filepath.Dir(filepath.Dir(filepath.Dir(filename)))
}

var projectRoot = rootDir()

func SqliteTestFile() string {
	return filepath.Join(projectRoot, "test.db")

}

func Favicon() string {
	return filepath.Join(projectRoot, "favicon.ico")

}

func SqliteProdFile() string {
	return filepath.Join(projectRoot, "prod.db")
}

func BooksJsonFile() string {
	return filepath.Join(projectRoot, "books.json")
}
