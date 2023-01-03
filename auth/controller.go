package auth

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"pawdot.app/utils"
)

var v validator.ValidationErrors

type Controller interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx)
}

type controller struct {
	us Service
	v  utils.Validator
}

// Register implements Controller
func (c *controller) Register(ctx *fiber.Ctx) error {
	var data IRegister

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errs := c.v.ValidateBody(data); errs != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "there was an error processing this request",
			"data":    errs,
		})
	}

	user, _ := c.us.FindByEmail(data.Email)
	if user != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "user with same email already exists",
		})
	}

	user, _ = c.us.FindByUsername(data.Username)
	if user != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "user with same Username already exists",
		})
	}

	userAuthData, err := c.us.CreateUser(data, ctx.IP())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "account could not be created",
			"data":    err,
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "account created successfully",
		"data":    userAuthData,
	})
}

// Login implements Controller
func (c *controller) Login(ctx *fiber.Ctx) {
	panic("unimplemented")
}

func InitController(us Service) Controller {
	return &controller{
		us: us,
	}
}
