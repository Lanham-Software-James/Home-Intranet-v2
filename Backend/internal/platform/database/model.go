// Package database servers as the wrapper to our Mongo DB Driver
package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model is the basic values for all records stored in our MongoDB
type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
