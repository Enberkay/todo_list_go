package main

import (
	"go_todo/config"
	"go_todo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDatabase()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
