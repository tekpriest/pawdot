package user

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"pawdot.app/models"
	"pawdot.app/utils"
)

type Service interface {
	CreateUser(data ICreateUser) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindOne(id string) (*models.User, error)
	UpdateUser(data interface{}) (*models.User, error)
	DebitWallet(userID string, amount float32) error
	CreditWallet(userID string, amount float32) error
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

	if err := s.db.Preload(clause.Associations).First(&user, "id = ?", id).Error; err != nil {
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

// CreditWallet implements Service
func (s *service) CreditWallet(userID string, amount float32) error {
	var wallet models.Wallet

	// if err := s.db.First(&wallet, "user_id = ?", userID).Error; err != nil {
	// 	if err.Error() == "record not found" {
	// 		go s.createWallet(userID, &wallet)
	// 	} else {
	// 		return err
	// 	}
	// }

	if err := s.db.Where(models.Wallet{UserID: userID}).Attrs(models.Wallet{
		Balance: wallet.Balance + amount,
	}).FirstOrCreate(&wallet).Error; err != nil {
		return err
	}

	return nil
}

// DebitWallet implements Service
func (s *service) DebitWallet(userID string, amount float32) error {
	var wallet models.Wallet

	if err := s.db.First(&wallet, "user_id = ?", userID).Error; err != nil {
		return err
	}

	if err := s.db.
		Table("wallets").
		Where("id =?", wallet.ID).
		Update("balance", wallet.Balance-amount).Error; err != nil {
		return err
	}

	return nil
}
