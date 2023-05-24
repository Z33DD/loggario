package control

import (
	"github.com/Z33DD/loggario/internal"
	"github.com/Z33DD/loggario/models"
	"github.com/gofiber/fiber/v2"
)

func AddLine(c *fiber.Ctx) error {
	line := new(models.Line)
	if err := c.BodyParser(&line); err != nil {
		return err
	}
	internal.DBConn.Create(&line)

	return c.JSON(line)
}

func GetLine(c *fiber.Ctx) error {
	lines := []models.Line{}

	internal.DBConn.First(&lines, c.Params("id"))

	return c.Status(200).JSON(lines)
}

func AllLines(c *fiber.Ctx) error {
	lines := []models.Line{}

	internal.DBConn.Find(&lines)

	return c.Status(200).JSON(lines)
}
