package main

import (
	"books/internal/database"
	"books/internal/paths"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"serde"
)

func main() {

	dburl := flag.String("db", "", "prod or test")
	var pathToDb string

	flag.Parse()

	switch *dburl {
	case "prod":
		pathToDb = paths.SqliteProdFile()
	case "test":
		pathToDb = paths.SqliteTestFile()
	default:
		log.Fatal("test or db")
	}

	db, err := sql.Open("sqlite3", pathToDb)

	queries := database.New(db)

	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatalf("os.Stat: %v", err)
	}

	bookModels, err := serde.DecodeJsonFileToStruct[[]database.BookModel](paths.BooksJsonFile())
	if err != nil {
		log.Fatalf("DecodeJsonFileToStruct to list of models %v", err)
	}

	for _, bm := range bookModels {
		addParams, err := bm.ParseAndValidate()
		if err != nil {
			log.Printf("parsing/validation error for book %v with error: %v", bm, err)
		}
		err = queries.AddBook(context.Background(), addParams)
		if err != nil {
			log.Printf("AddBook erroed %v with error %v", bm, err)
		}
	}

	allBooks, err := queries.GetAllBooks(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("All Books updated %+v\n", allBooks)

}
