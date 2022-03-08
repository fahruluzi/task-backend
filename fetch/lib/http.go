package lib

import "github.com/gofiber/fiber/v2"

// ResponseFormatter returning formatted JSON response
func ResponseFormatter(ctx *fiber.Ctx, code int, message string, body interface{}, err error) error {
	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   err.Error(),
		}
	} else {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   nil,
		}
	}

	return ctx.Status(code).JSON(response)
}
