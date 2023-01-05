package auth

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v9"
	"golang.org/x/crypto/bcrypt"
	"pawdot.app/user"
)

var ctx = context.Background()

type Service interface {
	CreateUser(data IRegister) (*UserAuthData, error)
	CheckIfUserExists(email, username string) []string
	createToken(data ICreateToken) (string, error)
	VerifyLogin(data IVerifyLogin) (*UserAuthData, error)
}

type service struct {
	us  user.Service
	rdb *redis.Client
}

// CreateUser implements Service
func (s *service) CreateUser(data IRegister) (*UserAuthData, error) {
	user, err := s.us.CreateUser(user.ICreateUser{
		Username:    data.Username,
		Email:       strings.ToLower(data.Email),
		Password:    hashPassword(data.Password),
		AccountType: data.AccountType,
	})
	if err != nil {
		return nil, err
	}
	token, err := s.createToken(ICreateToken{
		UserID:      user.ID,
		AccountType: user.AccountType,
	})
	if err != nil {
		return nil, err
	}

	return &UserAuthData{
		Token: token,
		User:  user,
	}, nil
}

// CheckIfUserExists implements Service
func (s *service) CheckIfUserExists(email, username string) []string {
	errors := make([]string, 0, 2)
	user, _ := s.us.FindByEmail(email)
	if user != nil {
		errors = append(errors, "user with same email already exists")
	}
	user, _ = s.us.FindByUsername(username)
	if user != nil {
		errors = append(errors, "user with same username already exists")
	}

	return errors
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func (s *service) createToken(data ICreateToken) (string, error) {
	defer ctx.Done()
	secret := os.Getenv("JWT_SECRET")
	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":      data.UserID,
		"accountType": data.AccountType,
		"expires":     time.Now().Add(time.Minute * 24 * 7).Unix(),
	})
	token, err := tokenData.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	// newToken := &models.AuthToken{
	// 	Token:  token,
	// 	IP:     data.IP,
	// 	UserID: data.UserID,
	// }

	return token, nil
}

func (s *service) VerifyLogin(data IVerifyLogin) (*UserAuthData, error) {
	user, err := s.us.FindByUsername(data.Username)
	if err != nil {
		return nil, err
	}
	if valid := validatePassword(data.Password, user.Password); !valid {
		return nil, errors.New("username and password does not match")
	}

	token, err := s.createToken(ICreateToken{
		UserID:      user.ID,
		AccountType: user.AccountType,
	})
	if err != nil {
		return nil, err
	}

	return &UserAuthData{
		Token: token,
		User:  user,
	}, nil
}

func validatePassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func InitService(us user.Service) Service {
	return &service{
		us: us,
	}
}
