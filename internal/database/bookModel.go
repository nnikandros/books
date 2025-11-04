package database

import "time"

type BookModel struct {
	Title           string
	Author          string
	PublicationDate time.Time
	FinishedDate    time.Time
	Rating          string
}

func (u BookModel) ToAddBookParams() AddBookParams {
	return AddBookParams(u)
}
