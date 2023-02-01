package auth

type RegistrationSuccessfulResponse struct {
	Success bool   `example:"true"`
	Mesage  string `example:"account created successfully"`
	Data    UserAuthData
} //	@Name	RegistrationSuccessfulResponse

type LoginSuccessfulResponse struct {
	Success bool   `example:"true"`
	Mesage  string `example:"account created successfully"`
	Data    UserAuthData
} //	@Name	LoginSuccessfulResponse
