package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time
}

func CreateUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	user.ID = uuid.New()
	DB.Table("users.users").Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User

	DB.Table("users.users").First(&user, "id = ?", id)

	return c.JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []User

	DB.Table("users.users").Find(&users)

	return c.JSON(users)	
}