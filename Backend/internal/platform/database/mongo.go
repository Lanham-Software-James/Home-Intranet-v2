// Package database servers as the wrapper to our Mongo DB Driver
package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect is used to create a new connection to our MongoDB
func Connect() (*mongo.Client, error) {
	username := os.Getenv("DB_USERNAME")
	password := url.QueryEscape(os.Getenv("DB_PASSWORD"))
	host := os.Getenv("DB_HOST")

	uri := fmt.Sprintf(`mongodb://%s:%s@%s:27017/?retryWrites=true&w=majority`, username, password, host)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client, nil
}

// Create is used to insert a new document into a collection
func (m *Model) Create(ctx context.Context, db *mongo.Database, collectionName string, model interface{}) error {
	collection := db.Collection(collectionName)

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	res, err := collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	m.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// Read is used to find one document based on a filter
func (m *Model) Read(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

// List is used to list all documents in a collection
func (m *Model) List(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, offset int64, limit int64, sort interface{}) ([]primitive.D, error) {
	opts := options.Find()
	opts.SetSkip(offset)
	opts.SetLimit(limit)
	opts.SetSort(sort)

	collection := db.Collection(collectionName)

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []primitive.D
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// Update is used to update a document in specified collection
func (m *Model) Update(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, update interface{}) error {
	collection := db.Collection(collectionName)

	m.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// Delete is used to delete a document in specified collection
func (m *Model) Delete(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
