package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (config MongoConfig) deleteDocument(client *mongo.Client) error {
	collection := client.Database(config.Database).Collection(config.Collection)
	_, err := collection.DeleteOne(context.TODO(), config.Data)

	if err != nil {
		return err
	}

	return nil
}

// DeleteHandler returns the delete hbs template located in the views folder
func DeleteHandler(c *fiber.Ctx) error {
	p := new(Person)

	if err := c.BodyParser(p); err != nil {
		return c.JSON(Response{
			ErrorMessage: "Formatting error",
			StatusCode:   400,
			Success:      false,
			Data:         nil,
		})
	}

	// struct can be found in db.go
	var mongoConfig MongoConfig = MongoConfig{
		Database:   "gofiber", // Which database
		Collection: "people",  // Which collection
		Data: bson.M{ // The data that will be inserted into the collection
			"_id": p.ID,
		},
	}

	mongoClient, err := connect() // temporarily connect to the database
	err = mongoConfig.deleteDocument(mongoClient)

	if err != nil {
		return c.JSON(Response{
			ErrorMessage: "MongoDB connection error",
			StatusCode:   400,
			Success:      false,
			Data:         nil,
		})
	}

	return c.JSON(Response{
		ErrorMessage: "Deleted success",
		StatusCode:   202,
		Success:      true,
		Data:         nil,
	})
}
