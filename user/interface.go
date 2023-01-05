package user

import "pawdot.app/models"

type ICreateUser struct {
	Username    string             `json:"username" validate:"required"`
	Email       string             `json:"email" validate:"required,email"`
	Password    string             `json:"password" validate:"required"`
	AccountType models.AccountType `json:"accountType" enum:"BUYER,SELLER"`
}
