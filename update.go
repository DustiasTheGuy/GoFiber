package main

import "github.com/gofiber/fiber/v2"

// GetUpdateHandler returns the update hbs template located in the views folder
func GetUpdateHandler(c *fiber.Ctx) error {
	return c.Render("update", fiber.Map{
		"title": "Update", // key value pairs
	}, "layouts/main") // Extend the layouts.hbs file
}
