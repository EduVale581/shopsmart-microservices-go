package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time
}

func CreateOrder(c *fiber.Ctx) error {
	order := new(Order)
	c.BodyParser(order)

	order.ID = uuid.New()
	order.Status = "CREATED"

	DB.Table("orders.orders").Create(&order)

	return c.JSON(order)
}

func GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order Order

	DB.Table("orders.orders").First(&order, "id = ?", id)

	return c.JSON(order)
}