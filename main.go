package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func index(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"title": "Test",
	})
}

func main() {
	//engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Static("/public", "./public")

	app.Get("/", index)

	app.Listen(":3000")
}
