// Package repository servers as the wrapper for our data persistance packages
package repository

import "go.mongodb.org/mongo-driver/mongo"

// Repository is the collection of data peristance wrappers
type Repository struct {
	Mongo *mongo.Database
}
