package main

import (
	"fmt"
	"github/golang/fiber/multidb/database"
	"github/golang/fiber/multidb/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("sdasd")
	db := database.New()
	productRrepo := repository.NewProduct(db)
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("tanant1", "1234") // Stores the string "admin" under a non-exported type key
		return c.Next()
	})
	app.Get("/", func(c *fiber.Ctx) error {
		t, ok := c.Locals("tanant1").(string) // Retrieves the data stored under the key and performs a type assertion

		return c.SendString(productRrepo.GetById())
	})
	app.Listen(":3000")
}
