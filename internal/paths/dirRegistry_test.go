package paths

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestDirectories(t *testing.T) {
	output := rootDir()

	one := filepath.Base(output)

	if one != "books" {
		t.Error("base is not books")
	}

	faviconPath := Favicon()

	if filepath.Base(faviconPath) != "favicon.ico" {
		t.Error("favicon.ico not matching the path that it should")
	}

	// p, _ := filepath.Abs(".")
	// fmt.Println(filepath.Dir(filepath.Dir(p)))

	fmt.Println(BooksJsonFile())
}
