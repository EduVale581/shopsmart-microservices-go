package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	InitDB()

	app := fiber.New()

	app.Post("/users", CreateUser)
	app.Get("/users/:id", GetUser)
	app.Get("/users", GetAllUsers)

	app.Listen(":3001")
}