package main

import (
	"context"

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

	return client, err
}

func createDocument(config MongoConfig, client *mongo.Client) error {
	collection := client.Database(config.Database).Collection(config.Collection)

	_, err := collection.InsertOne(context.TODO(), config.Data)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	//fmt.Println("Inserted a single document: ", result.InsertedID)

	return nil
}

func readDocuments(config MongoConfig, client *mongo.Client) ([]*Person, error) {
	collection := client.Database(config.Database).Collection(config.Collection)

	var result []*Person

	cur, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Person
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		result = append(result, &elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return result, nil
}
