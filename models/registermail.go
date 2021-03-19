package models

import (
	"time"

	"gorm.io/gorm"
)

type RegistEmail struct {
	ID          uint           `gorm:"primarykey, autoIncrement" json:"id"`
	Type        string         `json:"type"`
	Email       string         `json:"email"`
	Message     string         `json:"message"`
	Status      bool           `json:"status"`
	DeliveredAt time.Time      `json:"delivered_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
