package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p Person) createDocument(client *mongo.Client) error {
	collection := client.Database("gofiber").Collection("people")

	_, err := collection.InsertOne(context.TODO(), bson.M{ // The data that will be inserted into the collection
		"Name":       p.Name,
		"Department": p.Department,
		"Age":        p.Age,
		"Salary":     p.Salary,
	})
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())
	return nil
}

// CreateHandler for creating new documents
func CreateHandler(c *fiber.Ctx) error {
	p := new(Person)

	if err := c.BodyParser(p); err != nil {
		fmt.Println(err)
		return err
	}

	// Response struct can be found in main.go
	if len(p.Name) < 1 {
		return c.JSON(Response{
			ErrorMessage: "Name field too short",
			StatusCode:   406,
			Success:      false,
			Data:         nil,
		})

	} else if len(p.Department) < 1 {
		return c.JSON(Response{
			ErrorMessage: "Department field too short",
			StatusCode:   406,
			Success:      false,
			Data:         nil,
		})

	} else if p.Age <= 0 {
		return c.JSON(Response{
			ErrorMessage: "Age too low",
			StatusCode:   406,
			Success:      false,
			Data:         nil,
		})
	}

	mongoClient, err := connect() // temporarily connect to the database

	if err != nil { // The connection to the database was not successful
		fmt.Println(err)
		return c.JSON(Response{
			ErrorMessage: "Connection to database could not be established",
			StatusCode:   400,
			Success:      false,
			Data:         nil,
		})
	}

	err = p.createDocument(mongoClient)

	if err != nil { // the document insertion was not successfull
		fmt.Println(err)
		return c.JSON(Response{
			ErrorMessage: "Document has not been inserted",
			StatusCode:   400,
			Success:      false,
			Data:         nil,
		})
	}

	return c.JSON(Response{ // insertion success
		ErrorMessage: "Document Created",
		StatusCode:   201,
		Success:      true,
		Data:         nil,
	})
}
