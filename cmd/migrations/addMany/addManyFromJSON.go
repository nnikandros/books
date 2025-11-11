package main

import (
	"books/internal/database"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"serde"
)

func main() {

	file := flag.String("file", "", "path to the JSON file")
	dburl := flag.String("db", "", "path to the JSON file")
	flag.Parse()

	db, err := sql.Open("sqlite3", *dburl)

	queries := database.New(db)

	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Stat(*file)
	if err != nil {
		log.Fatalf("os.Stat: %v", err)
	}

	bookModels, err := serde.DecodeJsonFileToStruct[[]database.BookModel](*file)
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
	fmt.Println(allBooks)

}
