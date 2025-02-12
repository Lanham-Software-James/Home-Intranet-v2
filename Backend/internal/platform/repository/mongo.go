// Package repository servers as the wrapper for our data persistance packages
package repository

import (
	"Home-Intranet-v2-Backend/internal/platform/config"
	"Home-Intranet-v2-Backend/internal/platform/pluralizer"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect is used to create a new connection to our MongoDB
func Connect() (*mongo.Database, error) {
	username := config.GetDBUserName()
	password := config.GetDBPassword()
	host := config.GetDBHost()
	database := config.GetDBName()

	uri := fmt.Sprintf("mongodb://%s:%s@%s", username, password, host)

	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("issue connecting to Mongo Client: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("issue verifying Mongo Client connection: %w", err)
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
func (db *Repository) Read(ctx context.Context, model interface{}, filter interface{}) error {
	collectionName, err := getCollectionName(model)
	if err != nil {
		return err
	}

	collection := db.Mongo.Collection(collectionName)

	if err = collection.FindOne(ctx, filter).Decode(model); err != nil {
		return err
	}

	return nil
}

// List is used to list all documents in a collection
func (db *Repository) List(ctx context.Context, model interface{}, filter map[string]string, sort map[string]string, offset int64, limit int64) ([]byte, error) {
	collectionName, err := getCollectionName(model)
	if err != nil {
		return nil, err
	}

	opts := options.Find()
	opts.SetSkip(offset)
	opts.SetLimit(limit)
	opts.SetSort(buildBSON(sort))

	collection := db.Mongo.Collection(collectionName)

	cursor, err := collection.Find(ctx, buildBSON(filter), opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var bsonResults []bson.M
	if err = cursor.All(ctx, &bsonResults); err != nil {
		return nil, err
	}

	return json.Marshal(bsonResults)
}

// Update is used to update a document in specified collection
func (db *Repository) Update(ctx context.Context, model interface{}, filter interface{}) error {
	collectionName, err := getCollectionName(model)
	if err != nil {
		return err
	}

	collection := db.Mongo.Collection(collectionName)

	setDefaultFields(model, false)

	_, err = collection.UpdateOne(ctx, filter, model)
	if err != nil {
		return err
	}

	return nil
}

// Delete is used to delete a document in specified collection
func (db *Repository) Delete(ctx context.Context, model interface{}, filter interface{}) error {
	collectionName, err := getCollectionName(model)
	if err != nil {
		return err
	}

	collection := db.Mongo.Collection(collectionName)

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

// IsNotFoundError verifies the type of error returning from a find query
func (db *Repository) IsNotFoundError(err error) bool {
	return errors.Is(err, mongo.ErrNoDocuments)
}

func buildBSON(data map[string]string) bson.D {
	doc := bson.D{}

	for key, value := range data {
		num, err := strconv.Atoi(value)
		if err == nil {
			doc = append(doc, bson.E{Key: key, Value: num})
		} else {
			doc = append(doc, bson.E{Key: key, Value: value})
		}
	}

	return doc
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
