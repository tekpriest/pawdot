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
	us Service
	v  utils.Validator
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

	user, _ := c.us.FindByEmail(data.Email)
	if user != nil {
		return responses.BadRequestResponse(ctx, errors.New("user with same email already exists"))
	}

	user, _ = c.us.FindByUsername(data.Username)
	if user != nil {
		return responses.BadRequestResponse(ctx, errors.New("user with same username already exists"))
	}

	user, err := c.us.CreateUser(data, ctx.IP())
	if err != nil {
		return responses.BadRequestResponse(ctx, errors.New("account could not be created"), err)
	}

	userData, err := c.us.CreateAuthData(*user)
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

	user, err := c.us.FindByUsername(data.Username)
	if err != nil {
		responses.BadRequestResponse(ctx, errors.New("user with same username already exists"))
	}

	result, passwordOk := c.us.VerifyLogin(data.Username, data.Password)
	if !passwordOk {
		responses.BadRequestResponse(ctx, errors.New(result), errors.New("wrong credentials entered"))
	}

	userData, err := c.us.CreateAuthData(*user)
	if err != nil {
		return responses.BadRequestResponse(ctx, errors.New("there was an error with this request"), err)
	}

	responses.OkResponse(ctx, "logged in successfully", userData)
	return
}

func InitController(us Service) Controller {
	return &controller{
		us: us,
	}
}
