package main

import "github.com/gofiber/fiber/v2"

// UpdateHandler returns the update hbs template located in the views folder
func UpdateHandler(c *fiber.Ctx) error {
	return c.JSON(Response{
		ErrorMessage: "",
		StatusCode:   200,
		Success:      true,
	})
}
