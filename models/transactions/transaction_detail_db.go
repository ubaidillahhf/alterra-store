package transactions

import (
	"time"

	"gorm.io/gorm"
)

type TransactionDetail struct {
	Id            int            `json:"id" gorm:"primaryKey;index"`
	TransactionId int            `json:"transaction_id" gorm:"index;not null"`
	Name          string         `json:"name" gorm:"not null"`
	CategoryId    int            `json:"category_id" gorm:"not null"`
	Sku           string         `json:"sku" gorm:"not null"`
	Price         float64        `json:"price" gorm:"not null"`
	Description   string         `json:"description"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
