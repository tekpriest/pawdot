package auth

import (
	"github.com/gofiber/fiber/v2"
	"pawdot.app/utils"
)

func Route(r fiber.Router) {
	r = r.Group("/auth")
	db := utils.InitDatabaseConnection()
	s := InitService(db)
	c := InitController(s)

	r.Post("/register", c.Register)
}
