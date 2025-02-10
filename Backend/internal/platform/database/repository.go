package database

import "go.mongodb.org/mongo-driver/mongo"

// Repository is the collection of data peristance wrappers
type Repository struct {
	Mongo *mongo.Database
}
