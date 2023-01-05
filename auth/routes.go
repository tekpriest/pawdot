package auth

import (
	"github.com/gofiber/fiber/v2"
	"pawdot.app/user"
	"pawdot.app/utils"
)

func Route(r fiber.Router) {
	r = r.Group("/auth")
	db := utils.InitDatabaseConnection()
	us := user.InitService(db)
	s := InitService(us)
	c := InitController(s)

	r.Post("/register", c.Register)
	r.Post("/login", c.Login)
}
