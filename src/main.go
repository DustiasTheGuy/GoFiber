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
	// handler returns an error
	app.Get("/", GetCreateHandler)
	app.Post("/create", PostCreateHandler)

	app.Get("/read", GetReadHandler)
	app.Get("/update", GetUpdateHandler)
	app.Get("/delete", GetDeleteHandler)

	app.Listen(":3000")
}
