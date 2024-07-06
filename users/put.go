package users

import (
	"github.com/bhyago/fiber-mongodb/db"
	"github.com/gofiber/fiber/v2"
)

func updateUser(c *fiber.Ctx) error {
	body := new(User)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	var result User

	err := db.UpdateById("users", c.Params("id"), body, &result)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    result,
	})
}
