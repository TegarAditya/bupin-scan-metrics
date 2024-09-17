package routes

import (
	"bupin-scan-metrics/controllers"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// SWAGGER
	app.Get("/docs/*", swagger.HandlerDefault)

	// API GROUP
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	// ScanMetric routes
	api.Get("/scanmetrics", controllers.GetScanMetrics)
	api.Get("/scanmetrics/:id", controllers.GetScanMetric)
	api.Post("/scanmetrics", controllers.CreateScanMetric)
	api.Put("/scanmetrics/:id", controllers.UpdateScanMetric)
	api.Delete("/scanmetrics/:id", controllers.DeleteScanMetric)
}
