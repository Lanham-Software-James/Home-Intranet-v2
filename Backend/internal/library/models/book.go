// Package models stores all of our models for the library module
package models

import (
	"Home-Intranet-v2-Backend/internal/platform/repository"
	"time"
)

// Book is the type for books in our library
type Book struct {
	repository.Model `bson:",inline" json:",inline"`
	Title            string    `bson:"title" json:"title"`
	Authors          []Author  `bson:"authors" json:"authors"`
	Shelf            string    `bson:"shelf" json:"shelf"`
	CheckedOut       bool      `bson:"checked_out" json:"checked_out"`
	CheckedOutBy     string    `bson:"checked_out_by" json:"checked_out_by"`
	CheckedOutTime   time.Time `bson:"checked_out_time" json:"checked_out_time"`
}
