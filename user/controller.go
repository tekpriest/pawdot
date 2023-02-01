package user

import (
	"errors"

	// "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"pawdot.app/responses"
	// "pawdot.app/utils"
)

// var v validator.ValidationErrors

type Controller interface {
	Profile(ctx *fiber.Ctx) error
}

type controller struct {
	s Service
	// v utils.Validator
}

func InitController(s Service) Controller {
	return &controller{
		s: s,
	}
}

// Profile godoc
//
//	@Tags		User
//	@ID			profile
//	@Security	ApiKey
//	@Router		/user/profile [get]
//	@Success	201	{object}	GetProfileSuccessfulResponse
func (c *controller) Profile(ctx *fiber.Ctx) error {
	userId := ctx.Get("userId")

	user, err := c.s.FindOne(userId)
	if err != nil {
		return responses.BadRequestResponse(
			ctx,
			errors.New("there was an error with this request"),
			err,
		)
	}

	return responses.OkResponse(ctx, "fetched user profile", user)
}
