// Package models stores all of our models for the library module
package models

import (
	"Home-Intranet-v2-Backend/internal/platform/repository"

	"go.mongodb.org/mongo-driver/bson"
)

// Author is the type for authors in our library
type Author struct {
	repository.Model `bson:",inline" json:",inline"`
	FirstName        string `bson:"first_name" json:"first_name"`
	MiddleName       string `bson:"middle_name" json:"middle_name"`
	LastName         string `bson:"last_name" json:"last_name"`
	Suffix           string `bson:"suffix" json:"suffix"`
}

// MarshalBSON is used when the author are embedded in a book object and marshalled into BSON
func (a Author) MarshalBSON() ([]byte, error) {
	return bson.Marshal(bson.M{
		"first_name":  a.FirstName,
		"middle_name": a.MiddleName,
		"last_name":   a.LastName,
		"suffix":      a.Suffix,
	})
}
