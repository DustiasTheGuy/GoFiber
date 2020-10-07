package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Person struct is how we can interface with data submitted by a form
type Person struct {
	Name    string `form:"name"`    // Name of an employee
	Company string `form:"company"` // Company of an employee
	Age     int16  `form:"age"`     // Age of an employee
}

// GetCreateHandler returns the create hbs template located in the views folder
func GetCreateHandler(c *fiber.Ctx) error { // Create
	return c.Render("create", fiber.Map{
		"title":   "Create", // Arbitrary
		"error":   c.Query("error"),
		"success": c.Query("success"),
	}, "layouts/main") // Extend the layouts.hbs file
}

// PostCreateHandler for parsing incomming form data
func PostCreateHandler(c *fiber.Ctx) error {
	p := new(Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	if len(p.Name) < 5 {
		fmt.Println("Name field missing")
		return c.Redirect("/?error=Name field missing") // display some kind of error message
	} else if len(p.Company) < 5 {
		fmt.Println("Company field missing")
		return c.Redirect("/?error=Company field missing") // display some kind of error message
	} else if p.Age <= 0 {
		fmt.Println("Age to low or missing")
		return c.Redirect("/?error=Age to low or missing") // display some kind of error message
	}

	var mongoConfig MongoConfig = MongoConfig{
		Database:   "gofiber",
		Collection: "people",
		Data: bson.M{
			"Name":    p.Name,
			"Company": p.Company,
			"Age":     p.Age,
		},
		channel: make(chan error),
	}

	go connect(mongoConfig) // insert document with a go routine

	err := <-mongoConfig.channel

	if err != nil {
		return c.Redirect("/") // display some kind of error message
	}

	// insertion success
	return c.Redirect("/?success=Document saved")
}
