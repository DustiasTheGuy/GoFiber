package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func readDocuments(config MongoConfig, client *mongo.Client) ([]*Person, error) {
	var result []*Person
	collection := client.Database(config.Database).Collection(config.Collection)

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

// ReadHandler returns the read hbs template located in the views folder
func ReadHandler(c *fiber.Ctx) error {
	// struct can be found in db.go
	var mongoConfig MongoConfig = MongoConfig{
		Database:   "gofiber", // Which database
		Collection: "people",  // Which collection
		Data:       nil,
	}

	mongoClient, err := connect() // temporarily connect to the database

	if err != nil {
		return c.JSON(Response{
			ErrorMessage: "MongoDB connection error",
			StatusCode:   400,
			Success:      false,
			Data:         nil,
		})
	}

	result, err := readDocuments(mongoConfig, mongoClient)

	if err != nil {
		return c.JSON(Response{
			ErrorMessage: "MongoDB query error",
			StatusCode:   400,
			Success:      false,
			Data:         nil,
		})
	}

	return c.JSON(Response{
		ErrorMessage: "Success, Here are the documents!",
		StatusCode:   202,
		Success:      true,
		Data:         result,
	})
}
