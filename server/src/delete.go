package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p Person) deleteDocument(client *mongo.Client) error {
	collection := client.Database("gofiber").Collection("people")
	_, err := collection.DeleteOne(context.TODO(), bson.M{ // The data that will be inserted into the collection
		"_id": p.ID,
	})

	if err != nil {
		return err
	}

	return nil
}

// DeleteHandler returns proccesses an incomming delete request
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

	mongoClient, err := connect() // temporarily connect to the database
	err = p.deleteDocument(mongoClient)

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
