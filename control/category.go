package control

import (
	"strconv"

	"github.com/Z33DD/loggario/internal"
	"github.com/Z33DD/loggario/models"
	"github.com/gofiber/fiber/v2"
)

func AddCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(&category); err != nil {
		return err
	}
	internal.DBConn.Create(&category)

	return c.JSON(category)
}

func GetCategory(c *fiber.Ctx) error {
	categories := []models.Category{}
	internal.DBConn.First(&categories, c.Params("id"))

	return c.Status(200).JSON(categories)
}

// Update
func UpdateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	id, _ := strconv.Atoi(c.Params("id"))

	internal.DBConn.Model(&models.Category{}).Where("id = ?", id).Update("name", category.Name)

	return c.Status(200).JSON("updated")
}

// Delete
func DeleteCategory(c *fiber.Ctx) error {
	category := new(models.Category)

	id, _ := strconv.Atoi(c.Params("id"))

	internal.DBConn.Where("id = ?", id).Delete(&category)

	return c.Status(200).JSON("deleted")
}

func AllCategories(c *fiber.Ctx) error {
	categories := []models.Category{}
	internal.DBConn.Find(&categories)

	return c.Status(200).JSON(categories)
}
