package main

import (
	"books/internal/database"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"serde"
)

func main() {

	service := database.NewService()
	dir := flag.String("dir", "", "directory which has the json files")
	flag.Parse()

	fileInfo, err := os.Stat(*dir)
	if err != nil {
		log.Fatalf("os.Stat: %v", err)
	}

	if !fileInfo.IsDir() {
		log.Fatalf("%v that you pass as an argument for dir is not a dir", dir)
	}

	files, err := filepath.Glob(*dir + "/*.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)
	for _, f := range files {
		bm, err := serde.DecodeJsonFileToStruct[database.BookModel](f)
		if err != nil {
			fmt.Printf("filepath: %v erroed out %v", f, err)
		}

		addParams, err := bm.ParseAndValidate()
		if err != nil {
			fmt.Printf("filepath: %v parsing and validation error %v", f, err)

		}

		err = service.Queries.AddBook(context.Background(), addParams)
		if err != nil {
			fmt.Printf("filepath: %v adding book in the database %v", f, err)
		}

	}
	// service.Queries.AddBook(context.Background())

}
