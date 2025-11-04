package database

import "time"

type BookModel struct {
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationDate time.Time `json:"publication-date"`
	FinishedDate    time.Time `json:"finished-date"`
	Rating          string    `json:"rating"`
}

func (u BookModel) ToAddBookParams() AddBookParams {
	return AddBookParams(u)
}
