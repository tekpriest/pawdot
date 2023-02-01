package models

import (
	"time"

	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type (
	SaleType     string
	SalePriority int
	SaleStatus   string
)

const (
	NEW         SaleType = "NEW"
	PROMOTED    SaleType = "PROMOTED"
	REPUBLISHED SaleType = "REPUBLISHED"
)

const (
	LOW SalePriority = iota
	NORMAL
	HIGH
)

const (
	PENDING   SaleStatus = "PENDING"
	PUBLISHED SaleStatus = "PUBLISHED"
	CANCELLED SaleStatus = "CANCELLED"
	CLOSED    SaleStatus = "CLOSED"
)

type Sale struct {
	Base
	Type        SaleType     `json:"type"                gorm:"column:type;default:NEW"                      sql:"type:ENUM('NEW','PROMOTED','REPUBLISHED')"            example:"type"`
	Category    string       `json:"category"                                                                                                                           example:"DOG"`
	Priority    SalePriority `json:"priority"            gorm:"column:priority"`
	Breed       string       `json:"breed"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	StartingBid float32      `json:"startingBig"`
	TraderID    string       `json:"traderId"            gorm:"column:trader_id"`
	Status      SaleStatus   `json:"status"              gorm:"column:status;default:PENDING"                sql:"type:ENUM('PENDING','PUBLISHED','CANCELLED','CLOSED)"`
	Sold        bool         `json:"sold"`
	SoldTo      *string      `json:"soldTo,omitempty"    gorm:"column:sold_to"`
	BidCount    int64        `json:"bidCount"            gorm:"bid_count;->"`
	Bids        []Bid        `json:"bids,omitempty"      gorm:"foreignKey:sale_id;references:id;type:string"`
	ExpiresBy   *time.Time   `json:"expiresBy,omitempty" gorm:"column:expires_by"`
} //	@Name	Sale Model

func (s *Sale) BeforeCreate(d *gorm.DB) (err error) {
	s.ID = cuid.New()
	s.Priority = NORMAL
	s.ExpiresBy = nil
	s.Type = NEW
	s.Status = PENDING
	return
}

func (s *Sale) AfterUpdate(d *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	if s.Status == SaleStatus(REPUBLISHED) {
		s.Priority = LOW
	}
	return
}
