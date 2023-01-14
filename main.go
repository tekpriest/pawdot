package main

import (
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"pawdot.app/api"
	"pawdot.app/docs"
	"pawdot.app/utils"
)

// @title					Pawdot API Service
// @version				1.0
// @description			Pawdow API Service
// @license.name			MIT
// @securityDefinitions	apiKey ApiKey
// @in						header
// @name					x-auth-token
// @Accept					json
// @Produce				json
// @host					localhost:3000
// @BasePath				/api/
func main() {
	if err := godotenv.Load(); err != nil {
		utils.AppError(err, "an error occured while starting server")
	}

	api := api.InitRouter()
	docs.SwaggerInfo.Schemes = []string{"http"}
	api.Get("/docs/*", swagger.HandlerDefault)
	utils.AppError(api.InitServer())
}
