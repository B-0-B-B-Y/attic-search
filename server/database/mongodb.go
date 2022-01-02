package database

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Object : The database structure for each object stored in the database
type Object struct {
	ID          string   `json:"_id" bson:",omitempty"`
	Side        string   `json:"side,omitempty" bson:",omitempty"`
	Position    int      `json:"position,omitempty" bson:",omitempty"`
	Item        string   `json:"item,omitempty" bson:",omitempty"`
	Container   string   `json:"container,omitempty" bson:",omitempty"`
	Description string   `json:"description,omitempty" bson:",omitempty"`
	Frequent    bool     `json:"frequent,omitempty" bson:",omitempty"`
	Keywords    []string `json:"keywords,omitempty" bson:",omitempty"`
}

var client mongo.Client
var collection mongo.Collection
var searchableFields [4]string = [4]string{"item", "container", "description", "keywords"}

// init : Initialise the database handler
func init() {
	mongoClient, err := NewClient()
	if err != nil {
		return
	}

	client = *mongoClient
	collection = *client.Database("storage").Collection("items")
}

// Create a new MongoDB client
func NewClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
}

// Search for any items that contain the keyword specified in select fields and return any matches
func SearchForItem(keyword string) (Object, error) {
	var result Object
	var fieldsToSearch bson.A
	regex := primitive.Regex{Pattern: keyword, Options: "i"}

	for _, field := range searchableFields {
		fieldsToSearch = append(fieldsToSearch, bson.D{{
			Key: field, Value: regex,
		}})
	}

	filter := bson.D{
		{
			Key: "$or", Value: fieldsToSearch,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return result, errors.New("record does not exist")
	} else if err != nil {
		return result, err
	}

	return result, nil
}

// Inserts new items to the database, assuming they pass schema validation checks
func InsertNewItems(items []Object) error {
	var dataToInsert []interface{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, object := range items {
		dataToInsert = append(dataToInsert, object)
	}

	_, err := collection.InsertMany(ctx, dataToInsert)
	if err != nil {
		return err
	}

	return nil
}

// Update an existing record with supplied data
func UpdateExistingItem(data Object) error {
	id, err := primitive.ObjectIDFromHex(data.ID)
	if err != nil {
		return err
	}

	// Clear ID before writing to mongo, prevents double insertion of the same ID
	data.ID = ""

	filter := bson.M{"_id": id}
	doc, err := toDoc(data)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": doc,
	}
	result := collection.FindOneAndUpdate(context.Background(), filter, update)

	if result.Err() != nil {
		return result.Err()
	} else {
		return nil
	}
}

// Delete an item from the DB matching the object ID
func DeleteItem(objectId string) error {
	id, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

// Helper function to convert a struct object to a BSON document
func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
