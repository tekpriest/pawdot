package middlewares

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"pawdot.app/models"
	"pawdot.app/responses"
)

type IGetUserData struct {
	UserID      string
	AccountType models.AccountType
}

type Auth interface {
	VerifyJWTToken(ctx *fiber.Ctx) error
}

type auth struct{}

// VerifyJWTToken implements Auth
func (a *auth) VerifyJWTToken(ctx *fiber.Ctx) error {
	headerToken := ctx.GetReqHeaders()["X-Auth-Token"]
	if headerToken == "" {
		return responses.UnauthorizedResponse(ctx, errors.New("empty authorization token in header"))
	}

	userData, err := verifyToken(headerToken)
	if err != nil {
		return responses.UnauthorizedResponse(ctx, err, err)
	}

	ctx.Request().Header.Set("userId", userData.UserID)
	ctx.Request().Header.Set("accountType", string(userData.AccountType))

	return ctx.Next()
}

func verifyToken(s string) (*IGetUserData, error) {
	var userData IGetUserData

	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(s, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("invalid signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, fmt.Errorf("token is expired")
	}
	claims := token.Claims.(jwt.MapClaims)
	userId, _ := claims["userId"].(string)
	accountType, _ := claims["accountType"].(models.AccountType)
	userData.UserID = userId
	userData.AccountType = accountType

	return &userData, nil
}

func InitAuthMiddleware() Auth {
	return &auth{}
}
