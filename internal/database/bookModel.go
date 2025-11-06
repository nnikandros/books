package database

import "time"

type BookModel struct {
	Title        string `json:"title"`
	Author       string `json:"author"`
	FinishedDate string `json:"finished-date"`
	Rating       string `json:"rating"`
	UriThumbnail string `json:"uri-thumbnail"`
	Review       string `json:"review"`
}

// ADD VALIDATION
func (u BookModel) ParseAndValidate() (AddBookParams, error) {
	f, err := time.Parse(time.DateOnly, u.FinishedDate)
	if err != nil {
		return AddBookParams{}, err
	}

	return AddBookParams{Title: u.Title, Author: u.Author, FinishedDate: f, Rating: u.Rating, UriThumbnail: u.UriThumbnail, Review: u.Review}, nil
}
