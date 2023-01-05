package auth

import "pawdot.app/models"

type IRegister struct {
	Username    string             `json:"username" validate:"required"`
	Email       string             `json:"email" validate:"required,email"`
	Password    string             `json:"password" validate:"required"`
	AccountType models.AccountType `json:"accountType" enum:"BUYER,SELLER"`
} // @Name Register

type ILogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
} // @Name Login

type IRequestPasswordReset struct {
	Email string `json:"email" validate:"required,email"`
} // @Name Request Password Reset

type IResetPassword struct {
	Password string `json:"password" validate:"required"`
	Code     string `json:"code" validate:"required"`
} // @Name Reset Password

type ICreateToken struct {
	UserID      string
	AccountType models.AccountType
	IP          string
}

type UserAuthData struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

type IVerifyLogin struct {
	ILogin
	IP string
}
