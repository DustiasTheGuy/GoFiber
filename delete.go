package main

import "github.com/gofiber/fiber/v2"

// GetDeleteHandler returns the delete hbs template located in the views folder
func GetDeleteHandler(c *fiber.Ctx) error {
	return c.Render("delete", fiber.Map{
		"title": "Delete", // key value pairs
	}, "layouts/main") // Extend the layouts.hbs file
}
