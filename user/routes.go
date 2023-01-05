package user

import (
	"github.com/gofiber/fiber/v2"
	"pawdot.app/middlewares"
	"pawdot.app/utils"
)

func Route(r fiber.Router) {
	r = r.Group("/user")
	db := utils.InitDatabaseConnection()
	s := InitService(db)
	c := InitController(s)
	m := middlewares.InitAuthMiddleware()

	r.Get("/profile", m.VerifyJWTToken, c.Profile)
}
