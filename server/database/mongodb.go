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
	Side        string   `json:"side"`
	Position    int      `json:"position"`
	Item        string   `json:"item"`
	Container   string   `json:"container"`
	Description string   `json:"description"`
	Frequent    bool     `json:"frequent"`
	Keywords    []string `json:"keywords"`
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

// TODO - Implement inserting elements - should support 1 or many insertions from a single call
// func InsertNewItems(items []Object) {

// }
