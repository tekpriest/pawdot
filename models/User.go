package models

import (
	"database/sql/driver"
	"strings"

	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type AccountType string

const (
	BUYER  AccountType = "BUYER"
	SELLER AccountType = "SELLER"
)

type User struct {
	Base
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	ProfileImg  string      `json:"profileImg"`
	AccountType AccountType `json:"accountType" sql:"type:ENUM('BUYER','SELLER')" gorm:"column:type"`
} // @Name User

func (u *User) BeforeCreate(d *gorm.DB) (err error) {
	u.ID = cuid.New()
	u.Email = strings.ToLower(u.Email)
	return
}

func (ut *AccountType) Scan(value interface{}) (err error) {
	*ut = AccountType(value.([]byte))
	return
}

func (ut *AccountType) Value() (driver.Value, error) {
	return driver.String.ConvertValue(ut)
}
