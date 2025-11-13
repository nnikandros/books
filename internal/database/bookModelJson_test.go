package database

import (
	"books/internal/paths"
	"serde"
	"testing"
)

func TestBookModelJson(t *testing.T) {

	m, err := serde.DecodeJsonFileToStruct[[]BookModel](paths.BooksJsonFile())
	if err != nil {
		t.Error(err)
	}

	for _, bookModel := range m {
		_, err = bookModel.ParseAndValidate()
		if err != nil {
			t.Errorf("bookModel %v gave parsing error %v", bookModel, err)
		}
	}

}
