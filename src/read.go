package main

import (
	"github.com/gofiber/fiber/v2"
)

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
			ErrorMessage: "",
			StatusCode:   401,
			Success:      false,
			Data:         nil,
		})
	}

	result, err := readDocuments(mongoConfig, mongoClient)

	if err != nil {
		return c.JSON(Response{
			ErrorMessage: "",
			StatusCode:   401,
			Success:      false,
			Data:         nil,
		})
	}

	return c.JSON(Response{
		ErrorMessage: "",
		StatusCode:   200,
		Success:      true,
		Data:         result,
	})
}
