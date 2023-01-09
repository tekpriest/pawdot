package models

import (
	"time"

	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type Sale struct {
	Base
	Category    string     `json:"category"`
	Breed       string     `json:"breed"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StartingBid float32    `json:"startingBig"`
	TraderID    string     `json:"traderId"    gorm:"column:trader_id"`
	Sold        bool       `json:"sold"`
	SoldTo      *string    `json:"soldTo"      gorm:"column:sold_to"`
	Bids        []Bid      `json:"bids"        gorm:"foreignKey:sale_id;references:id;type:string"`
	ExpiresBy   *time.Time `json:"expiresBy"   gorm:"column:expires_by"`
}

func (s *Sale) BeforeCreate(d *gorm.DB) (err error) {
	s.ID = cuid.New()
	s.ExpiresBy = nil
	return
}

func (s *Sale) AfterUpdate(d *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
