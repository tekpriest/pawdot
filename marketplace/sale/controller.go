package sale

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"pawdot.app/responses"
	"pawdot.app/utils"
)

type Controller interface {
	CreateSale(ctx *fiber.Ctx) error
	GetAllSales(ctx *fiber.Ctx) error
	GetPersonalSales(ctx *fiber.Ctx) error
	GetPersonalBids(ctx *fiber.Ctx) error
	GetSale(ctx *fiber.Ctx) error
	CreateBid(ctx *fiber.Ctx) error
	GetSaleBids(ctx *fiber.Ctx) error
	CancelSale(ctx *fiber.Ctx) error
	RepublishSale(ctx *fiber.Ctx) error
}

type controller struct {
	s Service
	v utils.Validator
}

// CreateBid godoc
//
//	@Tags		Bid
//	@ID			creaateBid
//	@Security	ApiKey
//	@Router		/marketplace/sales/:saleID/bid [post]
//	@Param		saleID	path		string	true	"saleID"
//	@Success	200		{object}	CreateBidSuccessfulResponse
func (c *controller) CreateBid(ctx *fiber.Ctx) error {
	var data ICreateBid
	userID := ctx.Get("userId")
	saleID := ctx.Params("saleID")
	if saleID == "" {
		return responses.BadRequestResponse(ctx, errors.New("empty parameter"))
	}

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

	sale, err := c.s.PlaceBid(saleID, userID, data.Amount)
	if err != nil {
		if err.Error() == "record not found" {
			return responses.NotFoundResponse(ctx, errors.New("sale not found"))
		}
		return responses.BadRequestResponse(ctx, err, errors.New("sale not found"))
	}

	return responses.OkResponse(ctx, "bid placed successfully", sale)
}

// CreateSale godoc
//
//	@Tags		Sale
//	@ID			createSale
//	@Security	ApiKey
//	@Router		/marketplace/sales/new [post]
//	@Success	201	{object}	CreateSaleSuccessfulResponse
func (c *controller) CreateSale(ctx *fiber.Ctx) error {
	userID := ctx.Get("userId")
	var data ICreateSale

	if err := ctx.BodyParser(&data); err != nil {
		return responses.BadRequestResponse(ctx, err)
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

// GetAllSales godoc
//
//	@Tags		Sale
//	@ID			getAllSales
//	@Security	ApiKey
//	@Router		/marketplace/sales [get]
//
// @Param breed query string false "Breed"
// @Param category query string false "Category"
// @Param page query int false "Page" default:"1"
// @Param limit query int false "Limit" default:"20"
//
//	@Success	201	{object}	GetAllSalesSuccessfulResponse
func (c *controller) GetAllSales(ctx *fiber.Ctx) error {
	var query IQuerySales

	if err := ctx.QueryParser(&query); err != nil {
		return responses.BadRequestResponse(ctx, err)
	}

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

// GetSale godoc
//
//	@Tags		Sale
//	@ID			getSale
//	@Security	ApiKey
//	@Router		/marketplace/sales/:saleID [get]
//	@Param		saleID	path		string	true	"saleID"
//	@Success	201	{object}	GetSaleSuccessfulResponse
func (c *controller) GetSale(ctx *fiber.Ctx) error {
	saleID := ctx.Params("saleID")

	sale, err := c.s.FetchSale(saleID)
	if err != nil {
		return responses.NotFoundResponse(ctx, errors.New("sale not found"), err)
	}

	return responses.OkResponse(ctx, "fetched sale", sale)
}

// GetSaleBids godoc
//
//	@Tags		Bid
//	@ID			getSaleBids
//	@Security	ApiKey
//	@Router		/marketplace/sales/:saleID/bids [get]
//	@Param		saleID	path		string	true	"saleID"
//	@Success	201	{object}	GetSaleBidsSuccessfulResponse
func (c *controller) GetSaleBids(ctx *fiber.Ctx) error {
	saleID := ctx.Params("saleID")
	bids, err := c.s.FetchSaleBids(saleID)
	if err != nil {
		return responses.NotFoundResponse(ctx, errors.New("there was an issue fetching bids"))
	}
	return responses.OkResponse(ctx, "fetched all sale bids", bids)
}

// GetSaleBids godoc
//
//	@Tags		Bid
//	@ID			getPersonalBids
//	@Security	ApiKey
//	@Router		/marketplace/personal/bids [get]
//	@Success	201	{object}	GetPersonalBidsSuccessfulResponse
func (c *controller) GetPersonalBids(ctx *fiber.Ctx) error {
	userID := ctx.Get("userId")

	bids, err := c.s.FetchPersonalBids(userID)
	if err != nil {
		return responses.NotFoundResponse(ctx, errors.New("there was an issue fetching bids"))
	}
	return responses.OkResponse(ctx, "fetched all personal bids", bids)
}

// GetPersonalSales godoc
//
//	@Tags		Sale
//	@ID			getPersonalSales
//	@Security	ApiKey
//	@Router		/marketplace/personal/sales [get]
//
// @Param breed query string false "Breed"
// @Param category query string false "Category"
// @Param page query int false "Page" default:"1"
// @Param limit query int false "Limit" default:"20"
//
//	@Success	201	{object}	GetPersonalSalesSuccessfulResponse
func (c *controller) GetPersonalSales(ctx *fiber.Ctx) error {
	userID := ctx.Get("userId")
	var query IQuerySales

	if err := ctx.QueryParser(&query); err != nil {
		return responses.BadRequestResponse(ctx, err)
	}

	sales, err := c.s.FetchPersonalSales(userID, query)
	if err != nil {
		return responses.NotFoundResponse(ctx, errors.New("there was an issue fetching sales"))
	}

	return responses.OkResponse(ctx, "fetched all personal sales", sales)
}

// CancelSale godoc
//
//	@Tags		Sale
//	@ID			cancelSale
//	@Security	ApiKey
//	@Router		/marketplace/sales/:saleID/cancel [delete]
//
//	@Param		saleID	path		string	true	"saleID"
//
//	@Success	201	{object}	CancelSaleSuccessfulResposnse
func (c *controller) CancelSale(ctx *fiber.Ctx) error {
	saleID := ctx.Params("saleID")
	if saleID == "" {
		return responses.BadRequestResponse(ctx, errors.New("empty parameter"))
	}

	if err := c.s.CancelSale(saleID); err != nil {
		return responses.BadRequestResponse(ctx, err)
	}

	return responses.OkResponse(ctx, "sale cancelled successfully")
}

// RepublishSale godoc
//
//	@Tags		Sale
//	@ID			republishSale
//	@Security	ApiKey
//	@Router		/marketplace/sales/:saleID/republish [post]
//
//	@Param		saleID	path		string	true	"saleID"
//
//	@Success	201	{object}	CancelSaleSuccessfulResposnse
func (c *controller) RepublishSale(ctx *fiber.Ctx) error {
	saleID := ctx.Params("saleID")
	if saleID == "" {
		return responses.BadRequestResponse(ctx, errors.New("empty parameter"))
	}

	sale, err := c.s.RepublishSale(saleID)
	if err != nil {
		return responses.BadRequestResponse(ctx, err)
	}

	return responses.OkResponse(ctx, "sale republished successfully", sale)
}

func InitController(s Service) Controller {
	return &controller{s: s}
}
