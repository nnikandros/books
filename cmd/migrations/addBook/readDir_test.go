package main

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {

	matches, err := filepath.Glob("/ec/local/home/nikanni/my-programming/app-workspace/types/*.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(matches)

}
