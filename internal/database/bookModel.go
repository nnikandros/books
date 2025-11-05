package database

import "time"

// type BookModel struct {
// 	Title           string    `json:"title"`
// 	Author          string    `json:"author"`
// 	PublicationDate time.Time `json:"publication-date"`
// 	FinishedDate    time.Time `json:"finished-date"`
// 	Rating          string    `json:"rating"`
// }

type BookModel struct {
	Title        string `json:"title"`
	Author       string `json:"author"`
	FinishedDate string `json:"finished-date"`
	Rating       string `json:"rating"`
}

// func (u BookModel) ToAddBookParams() AddBookParams {
// 	return AddBookParams(u)
// }

func (u BookModel) ParseBookModel() (AddBookParams, error) {
	f, err := time.Parse(time.DateOnly, u.FinishedDate)
	if err != nil {
		return AddBookParams{}, err
	}

	return AddBookParams{Title: u.Title, Author: u.Author, FinishedDate: f, Rating: u.Rating}, nil
}

// experimental. prbalby delete
func FromBook(b Book) BookModel {

	return BookModel{Author: b.Author, Title: b.Title, FinishedDate: b.FinishedDate.Format(time.DateOnly), Rating: b.Rating}

}
