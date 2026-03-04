package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(Product)
	c.BodyParser(product)

	product.ID = uuid.New()
	DB.Table("inventory.products").Create(&product)

	return c.JSON(product)
}

func GetProducts(c *fiber.Ctx) error {
	var products []Product
	DB.Table("inventory.products").Find(&products)

	return c.JSON(products)
}