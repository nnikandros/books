package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
)

func TestCheckTableExists(t *testing.T) {

	db, err := sql.Open("sqlite3", "/ec/local/home/nikanni/my-programming/app-workspace/books/test.db")
	if err != nil {
		log.Fatal(err)
	}
	// row := db.QueryRowContext(context.Background(), "SELECT name FROM sqlite_master WHERE type='table' AND name='books'")
	row := db.QueryRowContext(context.Background(), "SELECT * FROM sqlite_master WHERE type='table' AND name='books'")
	if err != nil {
		log.Fatal(err)
	}
	var name string
	err = row.Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	// var name string
	// err = row.Scan(&name)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("❌ Table 'books' does not exist.")
	// } else if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("✅ Table 'books' exists.")
	// }
	// fmt.Printf("%+v", r)
	fmt.Println(name)

}
