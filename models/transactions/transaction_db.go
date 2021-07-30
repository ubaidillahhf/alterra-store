package transactions

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id         int            `json:"id" gorm:"primaryKey;index"`
	Number     string         `json:"number" gorm:"not null"`
	UserId     int            `json:"user_id" gorm:"not null"`
	Date       time.Time      `json:"date" gorm:"not null"`
	Status     string         `json:"status" gorm:"not null"`
	TotalPrice float64        `json:"total_price" gorm:"not null"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
