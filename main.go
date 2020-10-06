package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: handlebars.New("./views", ".hbs"),
	})

	// prefix static folder with /public
	app.Static("/public", "./public")

	// string, handler
	app.Get("/", CreateHandler)
	app.Get("/create", CreateHandler)
	app.Get("/read", ReadHandler)
	app.Get("/update", UpdateHandler)
	app.Get("/delete", DeleteHandler)
	app.Listen(":3000")
}

func CreateHandler(c *fiber.Ctx) error {
	// Create
	return c.Render("create", fiber.Map{
		"title": "Create", // Arbitrary
	}, "layouts/main") // Extend the layouts.hbs file
}

func ReadHandler(c *fiber.Ctx) error {
	// Read
	return c.Render("read", fiber.Map{
		"title": "Read", // Arbitrary
	}, "layouts/main") // Extend the layouts.hbs file
}

func UpdateHandler(c *fiber.Ctx) error {
	// Update
	return c.Render("update", fiber.Map{
		"title": "Update", // Arbitrary
	}, "layouts/main") // Extend the layouts.hbs file
}

func DeleteHandler(c *fiber.Ctx) error {
	// Delete
	return c.Render("delete", fiber.Map{
		"title": "Delete", // Arbitrary
	}, "layouts/main") // Extend the layouts.hbs file
}
