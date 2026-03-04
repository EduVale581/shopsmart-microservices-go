package main

import "github.com/gofiber/fiber/v2"

func main() {
	InitDB()

	app := fiber.New()

	app.Post("/orders", CreateOrder)
	app.Get("/orders/:id", GetOrder)

	app.Listen(":3003")
}