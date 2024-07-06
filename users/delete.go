package users

import (
	"github.com/bhyago/fiber-mongodb/db"
	"github.com/gofiber/fiber/v2"
)

func deleteUser(c *fiber.Ctx) error {
	err := db.DeleteById("users", c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).SendString("")
}
