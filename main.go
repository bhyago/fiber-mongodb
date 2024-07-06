package main

import (
	"github.com/bhyago/fiber-mongodb/users"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	v1 := app.Group("/v1")
	users.SetRoutes(v1)

	app.Listen(":3000")
}
