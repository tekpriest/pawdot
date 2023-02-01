package sale

import (
	"github.com/gofiber/fiber/v2"
	"pawdot.app/middlewares"
	"pawdot.app/utils"
)

func Route(r fiber.Router) {
	r = r.Group("/marketplace")
	db := utils.InitDatabaseConnection()
	s := InitService(db)
	c := InitController(s)
	m := middlewares.InitAuthMiddleware()

	// sales
	r.Get("/sales", m.VerifyJWTToken, c.GetAllSales)
	r.Get("/personal/sales", m.VerifyJWTToken, c.GetPersonalSales)
	r.Post("/sales/new", m.VerifyJWTToken, c.CreateSale)

	// bids
	r.Get("/personal/bids", m.VerifyJWTToken, c.GetPersonalBids)
	r.Get("/sales/bids/:saleID", m.VerifyJWTToken, c.GetSaleBids)
	r.Get("/sales/:saleID", m.VerifyJWTToken, c.GetSale)
	r.Post("/sales/:saleID/bid", m.VerifyJWTToken, c.CreateBid)
	r.Delete("/sales/:saleID/cancel", m.VerifyJWTToken, c.CancelSale)
	r.Post("/sales/:saleID/republish", m.VerifyJWTToken, c.RepublishSale)
}
