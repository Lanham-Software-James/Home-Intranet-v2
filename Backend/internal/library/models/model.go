// Package models stores all of our models for the library module
package models

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
