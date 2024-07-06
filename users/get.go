package users

import (
	"github.com/bhyago/fiber-mongodb/db"
	"github.com/gofiber/fiber/v2"
)

func getAll(c *fiber.Ctx) error {
	var documents []User

	err := db.Find("users", &documents)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(documents)
}

func getByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var document User

	err := db.FindById("users", id, &document)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(document)
}
