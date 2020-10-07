package main

import "github.com/gofiber/fiber/v2"

// GetReadHandler returns the read hbs template located in the views folder
func GetReadHandler(c *fiber.Ctx) error {
	return c.Render("read", fiber.Map{
		"title": "Read", // key value pairs
	}, "layouts/main") // Extend the layouts.hbs file
}
