package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig struct for inserting data into mongodb
type MongoConfig struct {
	Data bson.M
}

// Person struct is how we can interface with data submitted by a form
type Person struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id,omitempty"` // Id for the document
	Name       string             `json:"name"`                     // Name of an employee
	Department string             `json:"department"`               // Company of an employee
	Age        int                `json:"age"`                      // Age of an employee
	Salary     int                `json:"salary"`                   // Salary of an employee
}

func connect() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}
