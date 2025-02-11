// Package repository servers as the wrapper for our data persistance packages
package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model is the basic values for all records persisted
type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
