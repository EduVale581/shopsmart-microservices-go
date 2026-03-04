package main

import "github.com/gofiber/fiber/v2"

func main() {
	InitDB()

	app := fiber.New()

	app.Post("/products", CreateProduct)
	app.Get("/products", GetProducts)

	app.Listen(":3002")
}