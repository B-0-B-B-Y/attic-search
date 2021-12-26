package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client mongo.Client
var collection mongo.Collection

// init : Initialise the database handler
func init() {
	mongoClient, err := NewClient()
	if err != nil {
		return
	}

	client = *mongoClient
	collection = *client.Database("storage").Collection("items")
}

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

func SearchOneItem(keyword string) (float64, error) {
	var result struct {
		Value float64
	}
	filter := bson.D{primitive.E{Key: "name", Value: keyword}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("record does not exist")
	} else if err != nil {
		return 0, err
	}

	return result.Value, nil
}
