package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p Person) updateDocument(client *mongo.Client) error {
	collection := client.Database("gofiber").Collection("people")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": p.ID}, bson.M{
		"$set": bson.M{
			"Name":       p.Name,
			"Department": p.Department,
			"Age":        p.Age,
			"Salary":     p.Salary,
		},
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// UpdateHandler is used to update a single document
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

	err = p.updateDocument(mongoClient)

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
