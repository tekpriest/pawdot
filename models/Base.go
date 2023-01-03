package models

import (
	"fmt"
	"time"

	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string    `json:"id" readonly:"true"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;autoCreateTime;column:created_at" readonly:"true" example:"2022-08-21 21:08"`
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	timestamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04"))
	return []byte(timestamp), nil
}

func (b *Base) BeforeCreate(d *gorm.DB) (err error) {
	b.ID = cuid.New()
	return
}
