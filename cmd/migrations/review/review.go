package main

import (
	"books/internal/database"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	service := database.NewService()
	id := flag.String("id", "", "id in the  table to update")
	file := flag.String("file", "", "plaintext file")
	flag.Parse()

	absPath, err := filepath.Abs(*file)
	if err != nil {
		log.Fatal(err)
	}

	b, err := os.ReadFile(absPath)
	if err != nil {
		log.Fatalf("ReadFile %v", err)
	}

	o, err := strconv.Atoi(*id)
	if err != nil {
		log.Fatal(err)
	}
	d := database.UpdateReviewByIdParams{Review: string(b), ID: int64(o)}

	book, err := service.Queries.UpdateReviewById(context.Background(), d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", book)

}
