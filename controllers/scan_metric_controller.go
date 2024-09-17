package controllers

import (
	"bupin-scan-metrics/config"
	"bupin-scan-metrics/models"
	"bupin-scan-metrics/responses"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var scanMetricCollection *mongo.Collection = config.GetCollection(config.DB, "scanmetrics")

var validate = validator.New()

// Create a ScanMetric
// @Summary      Create a metric
// @Description  create a metric
// @Tags         metric
// @Accept       json
// @Produce      json
// @Param        scanMetric body models.ScanMetric true "ScanMetric object"
// @Success      201  {object}  models.ScanMetric
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /scanmetrics [post]
func CreateScanMetric(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var scanMetric models.ScanMetric

	// Parse JSON body
	if err := c.BodyParser(&scanMetric); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ScanMetricResponse{
			Status:  fiber.StatusBadRequest,
			Message: "error",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	// Validate the struct
	if validationErr := validate.Struct(&scanMetric); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ScanMetricResponse{
			Status:  fiber.StatusBadRequest,
			Message: "error",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}

	newScanMetric := models.ScanMetric{
		ID:        primitive.NewObjectID(),
		UserID:    scanMetric.UserID,
		DeviceID:  scanMetric.DeviceID,
		ContentID: scanMetric.ContentID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := scanMetricCollection.InsertOne(ctx, newScanMetric)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert scan metric",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newScanMetric)
}

// ShowAllMetrics godoc
// @Summary      Show all metrics
// @Description  get all metrics
// @Tags         metric
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.ScanMetric
// @Failure      500  {object}  error
// @Router       /scanmetrics [get]
func GetScanMetrics(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := scanMetricCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve scan metrics",
		})
	}

	var scanMetrics []models.ScanMetric
	if err = cursor.All(ctx, &scanMetrics); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse scan metrics",
		})
	}

	return c.JSON(scanMetrics)
}

// ShowAccount godoc
// @Summary      Show a metric
// @Description  get metric by ID
// @Tags         metric
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "scan metric ID"
// @Success      200  {object}  models.ScanMetric
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /scanmetrics/{id} [get]
func GetScanMetric(c *fiber.Ctx) error {
	scanMetricID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(scanMetricID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ScanMetric ID",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var scanMetric models.ScanMetric
	err = scanMetricCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&scanMetric)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "ScanMetric not found",
		})
	}

	return c.JSON(scanMetric)
}

// UpdateScanMetric godoc
// @Summary      Update a metric
// @Description  update a metric
// @Tags         metric
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "scan metric ID"
// @Param        scanMetric body models.ScanMetric true "ScanMetric object"
// @Success      200  {object}  models.ScanMetric
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /scanmetrics/{id} [put]
func UpdateScanMetric(c *fiber.Ctx) error {
	scanMetricID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(scanMetricID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ScanMetric ID",
		})
	}

	var scanMetric models.ScanMetric
	if err := c.BodyParser(&scanMetric); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": scanMetric,
	}
	_, err = scanMetricCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update scan metric",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ScanMetric updated successfully",
	})
}

// DeleteScanMetric godoc
// @Summary      Delete a metric
// @Description  delete a metric
// @Tags         metric
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "scan metric ID"
// @Success      200  {object}  string
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /scanmetrics/{id} [delete]
func DeleteScanMetric(c *fiber.Ctx) error {
	scanMetricID := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(scanMetricID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ScanMetric ID",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = scanMetricCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete scan metric",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ScanMetric deleted successfully",
	})
}
