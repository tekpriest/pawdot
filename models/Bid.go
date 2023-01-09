package models

import (
	"time"

	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type Bid struct {
	Base
	UserID string  `json:"userId,omitempty" gorm:"column:user_id"`
	Amount float32 `json:"amount"`
	SaleID string  `json:"saleId,omitempty" gorm:"column:sale_id"`
}

func (b *Bid) BeforeCreate(d *gorm.DB) (err error) {
	b.ID = cuid.New()
	return
}

func (b *Bid) AfterCreate(d *gorm.DB) (err error) {
	var count int64
	d.Table("bids").Count(&count)
	if count == 1 {
		d.Table("sales").
			Where("id =?", b.SaleID).
			Update("expires_by", time.Now().
				Add(time.Hour*24))
	}
	return
}
