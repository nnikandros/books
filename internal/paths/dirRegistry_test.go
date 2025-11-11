package paths

import (
	"path/filepath"
	"testing"
)

func TestDirectories(t *testing.T) {
	output := appDir()

	one := filepath.Base(output)

	if one != "books" {
		t.Error("base is not books")
	}

	faviconPath := Favicon()

	if filepath.Base(faviconPath) != "favicon.ico" {
		t.Error("favicon.ico not matching the path that it should")
	}
}
