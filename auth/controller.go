package auth

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"pawdot.app/responses"
	"pawdot.app/utils"
)

var v validator.ValidationErrors

type Controller interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type controller struct {
	s Service
	v utils.Validator
}

// Register implements Controller
func (c *controller) Register(ctx *fiber.Ctx) error {
	var data IRegister

	responses.BadRequestResponse(ctx, ctx.BodyParser(&data))

	responses.BadRequestResponse(
		ctx,
		errors.New("there was an error processing this request"),
		c.v.ValidateBody(data),
	)

	if errs := c.s.CheckIfUserExists(data.Email, data.Username); len(errs) > 0 {
		return responses.BadRequestResponse(ctx, errors.New("account could not be created"), errs)
	}

	userData, err := c.s.CreateUser(data)
	if err != nil {
		return responses.InternalServerErrorResponse(ctx, errors.New("there was an issue creating this user"), err)
	}

	if err != nil {
		return responses.BadRequestResponse(ctx, errors.New("there was an error with this request"), err)
	}

	return responses.CreatedResponse(ctx, "account created successfully", userData)
}

// Login implements Controller
func (c *controller) Login(ctx *fiber.Ctx) (err error) {
	var data ILogin

	responses.BadRequestResponse(ctx, ctx.BodyParser(&data))

	responses.BadRequestResponse(
		ctx,
		errors.New("there was an error processing this request"),
		c.v.ValidateBody(data),
	)

	userData, err := c.s.VerifyLogin(IVerifyLogin{
		ILogin: data,
		IP:     ctx.IP(),
	})
	if err != nil {
		responses.BadRequestResponse(ctx, err, errors.New("wrong credentials entered"))
		return
	}

	responses.OkResponse(ctx, "logged in successfully", userData)
	return
}

func InitController(s Service) Controller {
	return &controller{
		s: s,
	}
}
