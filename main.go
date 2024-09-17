package main

import (
	"bupin-scan-metrics/config"
	"bupin-scan-metrics/routes"

	"github.com/gofiber/fiber/v2"

	_ "bupin-scan-metrics/docs"
)

// @title Scan Metrics API
// @version 1.0
// @description This is an API for Scan Metrics

// @contact.name Tegar Aditya
// @contact.email dinopuguh@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func main() {
	app := fiber.New()

	config.ConnectDB()

	routes.SetupRoutes(app)

	app.Listen(":3030")
}
