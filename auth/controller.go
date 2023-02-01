package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"pawdot.app/responses"
	"pawdot.app/utils"
)

type Controller interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type controller struct {
	s Service
	v utils.Validator
}

// Register godoc
//	@Tags		Auth
//	@ID			register
//	@Router		/auth/register [post]
//	@Param		Register	body		IRegister	true	"Body"
//	@Success	201			{object}	RegistrationSuccessfulResponse
func (c *controller) Register(ctx *fiber.Ctx) error {
	var data IRegister

	if err := ctx.BodyParser(&data); err != nil {
		return responses.BadRequestResponse(ctx, err)
	}

	if err := c.v.ValidateBody(data); err != nil {
		return responses.BadRequestResponse(
			ctx,
			errors.New("there was an error processing this request"),
			err,
		)
	}

	if errs := c.s.CheckIfUserExists(data.Email, data.Username); len(errs) > 0 {
		return responses.BadRequestResponse(ctx, errors.New("account could not be created"), errs)
	}

	userData, err := c.s.CreateUser(data)
	if err != nil {
		return responses.InternalServerErrorResponse(
			ctx,
			errors.New("there was an issue creating this user"),
			err,
		)
	}

	if err != nil {
		return responses.BadRequestResponse(
			ctx,
			errors.New("there was an error with this request"),
			err,
		)
	}

	return responses.CreatedResponse(ctx, "account created successfully", userData)
}

// Login godoc
//	@Tags		Auth
//	@ID			login
//	@Router		/auth/login [post]
//	@Param		body	body		ILogin	true	"Body"
//	@Success	201		{object}	LoginSuccessfulResponse
func (c *controller) Login(ctx *fiber.Ctx) (err error) {
	var data ILogin

	if err := ctx.BodyParser(&data); err != nil {
		return responses.BadRequestResponse(ctx, err)
	}

	if err := c.v.ValidateBody(data); err != nil {
		return responses.BadRequestResponse(
			ctx,
			errors.New("there was an error processing this request"),
			err,
		)
	}
	// responses.BadRequestResponse(ctx, ctx.BodyParser(&data))
	//
	// responses.BadRequestResponse(
	// 	ctx,
	// 	errors.New("there was an error processing this request"),
	// 	c.v.ValidateBody(data),
	// )

	userData, err := c.s.VerifyLogin(IVerifyLogin{
		ILogin: data,
		IP:     ctx.IP(),
	})
	if err != nil {
		if err.Error() == "record not found" {
			return responses.NotFoundResponse(ctx, errors.New("account does not exist"), err)
		}
		return responses.BadRequestResponse(
			ctx,
			errors.New("wrong credentials entered"),
			err.Error(),
		)
	}

	return responses.OkResponse(ctx, "logged in successfully", userData)
}

func InitController(s Service) Controller {
	return &controller{
		s: s,
	}
}
