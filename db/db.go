package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func InitDb(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
			return nil,err

	}
	return client, nil
}

func InsertData(client *mongo.Client) {
	database := client.Database("mydb")
	collection := database.Collection("mycollection")

	// Create a BSON document to insert.
	data := bson.M{"name": "John Doe", "age": 30}

	// Insert the document into the collection.
	collection.InsertOne(context.TODO(), data)
}
