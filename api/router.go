package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"pawdot.app/auth"
	"pawdot.app/marketplace/sale"
	"pawdot.app/user"
)

type Router interface {
	fiber.Router
	InitServer() error
}

type router struct {
	*fiber.App
}

func InitRouter() Router {
	r := fiber.New()
	r.Use(logger.New())

	return &router{r}
}

func (r *router) InitServer() error {
	api := r.Group("/api")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"status":    true,
			"message":   "Server up and running",
			"timestamp": time.Now().String(),
		})
	})

	auth.Route(api)
	user.Route(api)
	sale.Route(api)

	return r.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
