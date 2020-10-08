package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Response for an incomming request
type Response struct {
	ErrorMessage string      `json:"message"`
	StatusCode   int         `json:"statusCode"`
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
}

func main() {
	app := fiber.New()

	// prefix static folder with /public
	// app.Static("/public", "./public")

	// string, handler
	// Default config
	app.Use(cors.New())

	app.Get("/read", ReadHandler)
	app.Post("/create", CreateHandler)
	app.Post("/delete", DeleteHandler)
	app.Post("/update", UpdateHandler)

	app.Listen(":3000")
}
