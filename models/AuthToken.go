package models

type AuthToken struct {
	Base
	Token  string `json:"token"`
	IP     string `json:"ip"`
	UserID string `json:"userId"`
} //	@Name	AuthToken
