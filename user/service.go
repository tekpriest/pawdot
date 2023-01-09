package user

import (
	"strings"

	"gorm.io/gorm"
	"pawdot.app/models"
	"pawdot.app/utils"
)

type Service interface {
	CreateUser(data ICreateUser) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindOne(id string) (*models.User, error)
	UpdateUser(data interface{}) (*models.User, error)
}

type service struct {
	db *gorm.DB
}

func InitService(db utils.DatabaseConnection) Service {
	return &service{db: db.GetDB()}
}

// CreateUser implements Service
func (s *service) CreateUser(data ICreateUser) (*models.User, error) {
	newUser := &models.User{
		Username:    data.Username,
		Email:       strings.ToLower(data.Email),
		Password:    data.Password,
		AccountType: data.AccountType,
	}
	if err := s.db.Create(newUser).Error; err != nil {
		return nil, err
	}

	return s.FindOne(newUser.ID)
}

// FindByEmail implements Service
func (s *service) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := s.db.First(&user, "email = ?", strings.ToLower(email)).Error; err != nil {
		return nil, err
	}

	return s.FindOne(user.ID)
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
func (s *service) UpdateUser(data interface{}) (*models.User, error) {
	var user models.User
	if err := s.db.Table("users").Updates(data).First(&user).Error; err != nil {
		return nil, err
	}

	return s.FindOne(user.ID)
}
