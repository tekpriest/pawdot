package models

import (
	"time"
)

type Base struct {
	ID        string    `json:"id,omitempty"        readonly:"true" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt,omitempty" readonly:"true" gorm:"type:timestamp;autoCreateTime;column:created_at" example:"2022-08-21 21:08"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"                 gorm:"type:timestamp;autoUpdateTime;column:updated_at"`
}
