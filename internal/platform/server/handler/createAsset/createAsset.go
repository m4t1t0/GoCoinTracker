package createAsset

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type requestBody struct {
	Asset    string `json:"asset" validate:"required,min=3,max=50,alphanum" msg:"'asset' must be alphanumeric and between 3 and 50 characters"`
	Interval int    `json:"interval" validate:"required,gt=0" msg:"'interval' must be a positive integer"`
}

// Handler processes POST /api/v1/assets requests.
// It expects a JSON body with fields: "asset" (string) and "interval" (integer).
func Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body requestBody

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "invalid request body",
				"message": "expected JSON with 'asset' as string and 'interval' as integer",
			})
		}

		if err := validate.Struct(body); err != nil {
			var errs validator.ValidationErrors
			errors.As(err, &errs)
			messages := make([]string, 0, len(errs))
			t := reflect.TypeOf(body)
			for _, e := range errs {
				if f, ok := t.FieldByName(e.StructField()); ok {
					if msg := f.Tag.Get("msg"); msg != "" {
						messages = append(messages, msg)
						continue
					}
				}
				messages = append(messages, e.Error())
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "validation error",
				"message": messages,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"asset":    body.Asset,
			"interval": body.Interval,
		})
	}
}
