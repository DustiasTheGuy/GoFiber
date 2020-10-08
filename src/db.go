package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig struct for inserting data into mongodb
type MongoConfig struct {
	Database   string
	Collection string
	Data       bson.M
}

func connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return nil, err
	}

	return client, err
}

func insertDocument(config MongoConfig, client *mongo.Client) (*mongo.InsertOneResult, error) {
	collection := client.Database(config.Database).Collection(config.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, config.Data)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	return result, nil
}
