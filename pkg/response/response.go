package response

import "github.com/gofiber/fiber/v2"

type Envelope struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Success(c *fiber.Ctx, status int, data interface{}) error {
	if status == fiber.StatusNoContent {
		return c.SendStatus(fiber.StatusNoContent)
	}
	return c.Status(status).JSON(Envelope{
		Code:    "OK",
		Message: "success",
		Data:    data,
	})
}

func Error(c *fiber.Ctx, status int, code, message string) error {
	return c.Status(status).JSON(Envelope{
		Code:    code,
		Message: message,
	})
}
