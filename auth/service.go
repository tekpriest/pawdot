package auth

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"pawdot.app/models"
	"pawdot.app/utils"
)

var ctx = context.Background()

type Service interface {
	CreateUser(data IRegister, ip string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindOne(id string) (*models.User, error)
	UpdateUser(data interface{}) (*models.User, error)
	CreateToken(data ICreateToken) (string, error)
	CreateAuthData(user models.User) (UserAuthData, error)
	VerifyLogin(username string, password string) (string, bool)
}

type service struct {
	db  *gorm.DB
	rdb *redis.Client
}

// CreateAuthData implements Service
func (s *service) CreateAuthData(user models.User) (UserAuthData, error) {
	var userData UserAuthData

	userData.User = &user

	token, err := s.CreateToken(ICreateToken{
		UserID:      user.ID,
		AccountType: user.AccountType,
	})
	if err != nil {
		return userData, err
	}
	userData.Token = token

	return userData, nil
}

// CreateUser implements Service
func (s *service) CreateUser(data IRegister, ip string) (*models.User, error) {
	newUser := &models.User{
		Username:    data.Username,
		Email:       strings.ToLower(data.Email),
		Password:    hashPassword(data.Password),
		AccountType: data.AccountType,
	}
	if err := s.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	user, err := s.FindOne(newUser.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByEmail implements Service
func (s *service) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := s.db.First(&user, "email = ?", strings.ToLower(email)).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByUsername implements Service
func (s *service) FindByUsername(username string) (*models.User, error) {
	var user models.User

	if err := s.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return s.FindOne(user.ID)
}

// FindOne implements Service
func (s *service) FindOne(id string) (*models.User, error) {
	var user models.User

	if err := s.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser implements Service
func (*service) UpdateUser(data interface{}) (*models.User, error) {
	panic("unimplemented")
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func (s *service) CreateToken(data ICreateToken) (string, error) {
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

func (s *service) VerifyLogin(username string, password string) (string, bool) {
	user, err := s.FindByUsername(username)
	if err != nil {
		return err.Error(), false
	}
	if valid := validatePassword(password, user.Password); !valid {
		return "username and password does not match", false
	}

	return user.ID, true
}

func validatePassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func InitService(db utils.DatabaseConnection) Service {
	return &service{
		db: db.GetDB(),
	}
}
