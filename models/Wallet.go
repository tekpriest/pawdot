package models

import (
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type Wallet struct {
	Base
	Balance float32 `json:"balance"`
	UserID  string  `gorm:"column:user_id;type:string"`
}

func (w *Wallet) BeforeCreate(d *gorm.DB) (err error) {
	w.ID = cuid.New()
	return
}
