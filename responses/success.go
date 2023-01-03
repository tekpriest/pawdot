package responses

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Success bool        `json:"success" default:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
} // @Name SuccessResponse

func CreatedResponse(c *fiber.Ctx, message string, data ...interface{}) error {
	return c.Status(http.StatusCreated).JSON(&SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func OkResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(http.StatusOK).JSON(&SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}
