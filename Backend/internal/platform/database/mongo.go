// Package database servers as the wrapper to our Mongo DB Driver
package database

import (
	"Home-Intranet-v2-Backend/internal/platform/pluralizer"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect is used to create a new connection to our MongoDB
func Connect() (*mongo.Database, error) {
	username := os.Getenv("DB_USERNAME")
	password := url.QueryEscape(os.Getenv("DB_PASSWORD"))
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")

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

	return client.Database(database), nil
}

// Create is used to insert a new document into a collection
func (db *Repository) Create(ctx context.Context, model interface{}) error {
	collectionName, err := getCollectionName(model)
	if err != nil {
		return err
	}

	collection := db.Mongo.Collection(collectionName)

	value := reflect.ValueOf(model)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return fmt.Errorf("model must be a struct or pointer to struct")
	}

	idField := value.FieldByName("ID")

	model, err = setDefaultFields(model, true)
	if err != nil {
		return err
	}

	res, err := collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	if idField.IsValid() && idField.CanSet() {
		idField.Set(reflect.ValueOf(res.InsertedID.(primitive.ObjectID)))
	}

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
func (db *Repository) List(ctx context.Context, model interface{}, filter interface{}, sort interface{}, offset int64, limit int64) ([]primitive.D, error) {
	collectionName, err := getCollectionName(model)
	if err != nil {
		return nil, err
	}

	opts := options.Find()
	opts.SetSkip(offset)
	opts.SetLimit(limit)
	opts.SetSort(sort)

	collection := db.Mongo.Collection(collectionName)

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

func getCollectionName(model interface{}) (string, error) {

	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		return "", fmt.Errorf("model not a pointer")
	}

	elemType := reflect.TypeOf(model).Elem()

	if elemType.Kind() != reflect.Struct {
		return "", fmt.Errorf("model not a struct")
	}

	modelName := reflect.TypeOf(model).Elem().Name()
	modelName = strings.ToLower(modelName)

	collectionName := pluralizer.ToPlural(modelName)

	return collectionName, nil
}

func setDefaultFields(model interface{}, setCreate bool) (interface{}, error) {

	now := time.Now().UTC()

	value := reflect.ValueOf(model)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("model must be a struct or pointer to struct")
	}

	if setCreate {

		createdAtField := value.FieldByName("CreatedAt")
		if createdAtField.IsValid() && createdAtField.CanSet() {
			createdAtField.Set(reflect.ValueOf(now))
		}
	}

	updatedAtField := value.FieldByName("UpdatedAt")
	if updatedAtField.IsValid() && updatedAtField.CanSet() {
		updatedAtField.Set(reflect.ValueOf(now))
	}

	return model, nil
}
