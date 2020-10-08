package main

import (
	"github.com/gofiber/fiber/v2"
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
	app.Static("/public", "./public")

	// string, handler
	// handler returns an error
	app.Post("/create", CreateHandler)
	app.Get("/read", ReadHandler)
	app.Get("/update", UpdateHandler)
	app.Get("/delete", DeleteHandler)

	app.Listen(":3000")
}
