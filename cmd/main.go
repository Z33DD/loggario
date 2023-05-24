package main

import (
	"log"

	"github.com/Z33DD/loggario/control"
	"github.com/Z33DD/loggario/internal"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "loggario", EnablePrintRoutes: false})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} \n${body}\n\n",
	}))

	app.Use(keyauth.New(keyauth.Config{
		KeyLookup: "header:API_KEY",
		Validator: internal.ValidateAPIKey,
	}))
	internal.ConnectDb()

	app.Post("/line/", control.AddLine)
	app.Get("/line/:id", control.GetLine)
	app.Get("/line/", control.AllLines)

	app.Post("/category/", control.AddCategory)

	log.Fatal(app.Listen(":3000"))
}
