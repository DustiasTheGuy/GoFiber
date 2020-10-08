package main

import "github.com/gofiber/fiber/v2"

// DeleteHandler returns the delete hbs template located in the views folder
func DeleteHandler(c *fiber.Ctx) error {
	return c.JSON(Response{
		ErrorMessage: "",
		StatusCode:   200,
		Success:      true,
	})
}
