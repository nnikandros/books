package main

import (
	"books/internal/database"
	"books/internal/paths"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	dburl := flag.String("db", "", "prod or test")
	path := flag.String("path", "", "path to the json file")
	flag.Parse()
	var pathToDb string

	switch *dburl {
	case "prod":
		pathToDb = paths.SqliteProdFile()
	case "test":
		pathToDb = paths.SqliteTestFile()
	default:
		log.Fatal("test or db")
	}

	db, err := sql.Open("sqlite3", pathToDb)
	if err != nil {
		log.Fatalf("at sql.Open %v", err)
	}

	queries := database.New(db)

	_, err = os.Stat(*path)
	if err != nil {
		log.Fatalf("os.Stat: %v", err)
	}

	f, err := os.Open(*path)
	if err != nil {
		log.Fatalf("at os.Open file %v, %v", *path, err)
	}

	err = addBook(queries, f)

	if err != nil {
		log.Fatalf("at addBook %v", err)
	}

}

func addBook(queries *database.Queries, r io.ReadCloser) error {

	addBookParams, err := database.NewAddBookParams(r)
	if err != nil {
		return fmt.Errorf("database.NewAddBookParams(%v), %w", r, err)
	}

	err = queries.AddBook(context.Background(), addBookParams)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil

}
