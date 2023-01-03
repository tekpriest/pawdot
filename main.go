package main

import (
	"github.com/joho/godotenv"
	"pawdot.app/api"
	"pawdot.app/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		utils.AppError(err, "an error occured while starting server")
	}

	api := api.InitRouter()
	utils.AppError(api.InitServer())
}
