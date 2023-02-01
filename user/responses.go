package user

import "pawdot.app/models"

type GetProfileSuccessfulResponse struct {
	Success bool   `example:"true"`
	Mesage  string `example:"fetched user profile"`
	Data    models.User
} //	@Name	GetProfileSuccessfulResponse
