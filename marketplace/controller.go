package marketplace

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"pawdot.app/responses"
	"pawdot.app/utils"
)

var v validator.ValidationErrors

type Controller interface {
	CreateSale(ctx *fiber.Ctx) error
	GetAllSales(ctx *fiber.Ctx) error
	GetPersonalSales(ctx *fiber.Ctx) error
	GetPersonalBids(ctx *fiber.Ctx) error
	GetSale(ctx *fiber.Ctx) error
	CreateBid(ctx *fiber.Ctx) error
	GetSaleBids(ctx *fiber.Ctx) error
}

type controller struct {
	s Service
	v utils.Validator
}

// CreateBid implements Controller
func (c *controller) CreateBid(ctx *fiber.Ctx) error {
	var data ICreateBid
	userID := ctx.Get("userId")
	saleID := ctx.Params("saleID")
	if saleID == "" {
		return responses.BadRequestResponse(ctx, errors.New("empty parameter"))
	}

	responses.BadRequestResponse(ctx, ctx.BodyParser(&data))

	responses.BadRequestResponse(
		ctx,
		errors.New("there was an error processing this request"),
		c.v.ValidateBody(data),
	)

	sale, err := c.s.PlaceBid(saleID, userID, data.Amount)
	fmt.Println(err)
	if err != nil {
		return responses.NotFoundResponse(ctx, errors.New("sale not found"), err)
	}

	return responses.OkResponse(ctx, "bid placed successfully", sale)
}

// CreateSale implements Controller
func (c *controller) CreateSale(ctx *fiber.Ctx) error {
	userID := ctx.Get("userId")
	var data ICreateSale

	if err := ctx.BodyParser(&data); err != nil {
		responses.BadRequestResponse(ctx, err)
	}

	if errs := c.v.ValidateBody(data); len(errs) > 0 {
		return responses.BadRequestResponse(
			ctx,
			errors.New("there was an error processing this request"),
			errs,
		)
	}

	sale, err := c.s.CreateSale(userID, data)
	if err != nil {
		return responses.BadRequestResponse(
			ctx,
			errors.New("there was an error creating this sale"),
			err,
		)
	}

	return responses.CreatedResponse(ctx, "sale created successfully", sale)
}

// GetAllSales implements Controller
func (c *controller) GetAllSales(ctx *fiber.Ctx) error {
	var query IQuerySales

	responses.BadRequestResponse(ctx, ctx.QueryParser(&query))

	if query.Limit == 0 {
		query.Limit = 20
	}
	if query.Page == 0 {
		query.Page = 1
	}

	sales, err := c.s.FetchAllSales(query)
	if err != nil {
		return responses.BadRequestResponse(
			ctx,
			errors.New("there was an issue with this request"),
			err,
		)
	}

	return responses.OkResponse(ctx, "fetched all current sales", sales)
}

// GetSale implements Controller
func (c *controller) GetSale(ctx *fiber.Ctx) error {
	saleID := ctx.Params("saleID")

	sale, err := c.s.FetchSale(saleID)
	if err != nil {
		responses.NotFoundResponse(ctx, errors.New("sale not found"), err)
	}

	return responses.OkResponse(ctx, "fetched sale", sale)
}

// GetSaleBids implements Controller
func (c *controller) GetSaleBids(ctx *fiber.Ctx) error {
	saleID := ctx.Params("saleID")
	bids, err := c.s.FetchSaleBids(saleID)
	if err != nil {
		return responses.NotFoundResponse(ctx, errors.New("there was an issue fetching bids"))
	}
	return responses.OkResponse(ctx, "fetched all sale bids", bids)
}

// GetPersonalBids implements Controller
func (c *controller) GetPersonalBids(ctx *fiber.Ctx) error {
	userID := ctx.Get("userId")

	bids, err := c.s.FetchPersonalBids(userID)
	if err != nil {
		return responses.NotFoundResponse(ctx, errors.New("there was an issue fetching bids"))
	}
	return responses.OkResponse(ctx, "fetched all personal bids", bids)
}

// GetPersonalSales implements Controller
func (c *controller) GetPersonalSales(ctx *fiber.Ctx) error {
	userID := ctx.Get("userId")
	var query IQuerySales

	responses.BadRequestResponse(ctx, ctx.QueryParser(&query))

	sales, err := c.s.FetchPersonalSales(userID, query)
	if err != nil {
		return responses.NotFoundResponse(ctx, errors.New("there was an issue fetching sales"))
	}

	return responses.OkResponse(ctx, "fetched all personal sales", sales)
}

func InitController(s Service) Controller {
	return &controller{s: s}
}
