package main

import "github.com/gofiber/fiber/v2"

// ReadHandler returns the read hbs template located in the views folder
func ReadHandler(c *fiber.Ctx) error {
	return c.JSON(Response{
		ErrorMessage: "",
		StatusCode:   200,
		Success:      true,
	})
}
