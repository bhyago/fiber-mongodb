package users

import (
	"github.com/bhyago/fiber-mongodb/db"
	"github.com/gofiber/fiber/v2"
)

func addUser(c *fiber.Ctx) error {
	body := new(User)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid json")
	}

	id, err := db.Insert("users", body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	body.ID = id

	return c.Status(fiber.StatusCreated).JSON(body)
}
