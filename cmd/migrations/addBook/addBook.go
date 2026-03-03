package main

import (
	"books/internal/database"
	"books/internal/paths"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"serde"
)

func main() {

	dburl := flag.String("db", "test", "prod or test")
	var pathToDb string

	flag.Parse()

	args := flag.Args()
	if len(args) > 1 {
		log.Fatal("please provide onle one")
	}

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
		log.Fatal(err)
	}
	queries := database.New(db)

	f, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	bookModel, err := serde.DecodeJsonFileToStructV2[database.BookModel](f)
	if err != nil {
		log.Fatal(err)
	}

	addParams, err := bookModel.ParseAndValidate()

	err = queries.AddBook(context.Background(), addParams)
	if err != nil {
		log.Fatal(err)
	}

	allBooks, err := queries.GetAllBooks(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("All Books updated %+v\n", allBooks)
}
