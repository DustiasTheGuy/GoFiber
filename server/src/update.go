package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (config MongoConfig) updateDocument(client *mongo.Client, update bson.M) error {
	collection := client.Database(config.Database).Collection(config.Collection)
	_, err := collection.UpdateOne(context.TODO(), config.Data, update)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// UpdateHandler returns the update hbs template located in the views folder
func UpdateHandler(c *fiber.Ctx) error {
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
	if err != nil {
		return c.JSON(Response{
			ErrorMessage: "MongoDB connection error",
			StatusCode:   400,
			Success:      false,
			Data:         nil,
		})
	}

	err = MongoConfig{
		Database:   "gofiber", // Which database
		Collection: "people",  // Which collection
		Data:       bson.M{"_id": p.ID},
	}.updateDocument(mongoClient,
		bson.M{
			"$set": bson.M{
				"Name":       p.Name,
				"Department": p.Department,
				"Age":        p.Age,
				"Salary":     p.Salary,
			},
		},
	)

	if err != nil {
		return c.JSON(Response{
			ErrorMessage: "Update error",
			StatusCode:   400,
			Success:      false,
			Data:         nil,
		})
	}

	return c.JSON(Response{
		ErrorMessage: "Update success",
		StatusCode:   200,
		Success:      true,
		Data:         nil,
	})
}
