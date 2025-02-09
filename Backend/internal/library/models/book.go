// Package models stores all of our models for the library module
package models

import "Home-Intranet-v2-Backend/internal/platform/database"

// Book is the type for books in our library
type Book struct {
	database.Model  `bson:",inline"`
	Title           string `bson:"title"`
	AuthorFirstName string `bson:"author_first_name"`
	AuthorLastName  string `bson:"author_last_name"`
}
