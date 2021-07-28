package transactions

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id            int            `json:"id" gorm:"primaryKey;index"`
	TransactionId int            `json:"transaction_id" gorm:"index;not null"`
	Name          string         `json:"name" gorm:"not null"`
	Status        string         `json:"status" gorm:"not null"`
	Description   string         `json:"description"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
