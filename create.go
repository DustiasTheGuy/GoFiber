package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
		"title": "Create", // Arbitrary
	}, "layouts/main") // Extend the layouts.hbs file
}

// PostCreateHandler for parsing incomming form data
func PostCreateHandler(c *fiber.Ctx) error {
	p := new(Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	fmt.Println(p)

	// save person(p) to database

	return c.Redirect("/")
}
