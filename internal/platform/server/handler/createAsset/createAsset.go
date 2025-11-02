package createAsset

import (
	"github.com/gofiber/fiber/v2"
)

// Handler processes POST /api/v1/assets requests.
// It expects a JSON body with fields: "asset" (string) and "interval" (integer).
func Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			Asset    string `json:"asset"`
			Interval int    `json:"interval"`
		}

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "invalid request body",
				"message": "expected JSON with 'asset' as string and 'interval' as integer",
			})
		}

		if err := validateCreateAssetRequest(body.Asset, body.Interval); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "validation error",
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"asset":    body.Asset,
			"interval": body.Interval,
		})
	}
}

// validateCreateAssetRequest validates the required fields for creating an asset.
func validateCreateAssetRequest(asset string, interval int) error {
	if asset == "" {
		return fiber.NewError(fiber.StatusBadRequest, "'asset' is required and must be a non-empty string")
	}
	if interval <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "'interval' is required and must be a positive integer")
	}
	return nil
}
