package database

import (
	"fmt"
	"path/filepath"
	"serde"
	"testing"
)

func TestBookModelJson(t *testing.T) {

	p := "../../books-json-files"
	absPath, err := filepath.Abs(p)
	if err != nil {
		t.Error(err)
	}

	filepath.Join(absPath, "gravity.json")

	// m, err := serde.DecodeJsonFileToStruct[BookModel](filepath.Join(absPath, "gravity.json"))
	// if err != nil {
	// 	t.Error(err)
	// }

	// fmt.Printf("%+v", m)
	// fmt.Println(m.ParseAndValidate())

	m, err := serde.DecodeJsonFileToStruct[[]BookModel](filepath.Join(absPath, "books.json"))
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v", m)
	for _, b := range m {
		bookPArams, err := b.ParseAndValidate()
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%+v", bookPArams)
	}

}
