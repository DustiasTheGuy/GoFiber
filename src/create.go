package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Person struct is how we can interface with data submitted by a form
type Person struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"` // Id for the document
	Name    string             `json:"name"`                     // Name of an employee
	Company string             `json:"company"`                  // Company of an employee
	Age     int16              `json:"age"`                      // Age of an employee
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

// CreateHandler for parsing incomming form data
func CreateHandler(c *fiber.Ctx) error {
	p := new(Person)

	fmt.Println(p)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	// Response struct can be found in main.go
	if len(p.Name) < 1 {
		return c.JSON(Response{
			ErrorMessage: "Name too short",
			StatusCode:   406,
			Success:      false,
			Data:         nil,
		})

	} else if len(p.Company) < 1 {
		return c.JSON(Response{
			ErrorMessage: "Comapny too short",
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

	// struct can be found in db.go
	var mongoConfig MongoConfig = MongoConfig{
		Database:   "gofiber", // Which database
		Collection: "people",  // Which collection
		Data: bson.M{ // The data that will be inserted into the collection
			"Name":    p.Name,
			"Company": p.Company,
			"Age":     p.Age,
		},
	}

	mongoClient, err := connect() // temporarily connect to the database

	if err != nil { // The connection to the database was not successful
		return c.JSON(Response{
			ErrorMessage: "Connection to database could not be established",
			StatusCode:   400,
			Success:      false,
			Data:         nil,
		})
	}

	err = createDocument(mongoConfig, mongoClient)

	if err != nil { // the document insertion was not successfull
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
