package responses

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Success bool        `json:"success"        default:"false"`
	Message string      `json:"message"`
	Error   interface{} `json:"data,omitempty"`
	Path    string      `json:"path"`
} // @Name ErrorResponse

func BadRequestResponse(c *fiber.Ctx, err error, data ...interface{}) error {
	var error interface{}

	if len(data) > 0 {
		error = data[0]
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&ErrorResponse{
			Success: false,
			Message: err.Error(),
			Error:   error,
			Path:    c.Path(),
		})
	}
	return nil
}

func InternalServerErrorResponse(c *fiber.Ctx, err error, data ...interface{}) error {
	var error interface{}

	if len(data) > 0 {
		error = data[0]
	}
	return c.Status(http.StatusInternalServerError).JSON(&ErrorResponse{
		Success: false,
		Message: err.Error(),
		Error:   error,
		Path:    c.Path(),
	})
}

func UnauthorizedResponse(c *fiber.Ctx, err error, data ...interface{}) error {
	var error interface{}

	if len(data) > 0 {
		error = data[0]
	}
	return c.Status(http.StatusInternalServerError).JSON(&ErrorResponse{
		Success: false,
		Message: err.Error(),
		Error:   error,
		Path:    c.Path(),
	})
}

func NotFoundResponse(c *fiber.Ctx, err error, data ...interface{}) error {
	var error interface{}

	if len(data) > 0 {
		error = data[0]
	}
	return c.Status(http.StatusNotFound).JSON(&ErrorResponse{
		Success: false,
		Message: err.Error(),
		Error:   error,
		Path:    c.Path(),
	})
}
