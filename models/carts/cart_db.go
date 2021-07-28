package carts

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	Id           int            `json:"id" gorm:"primaryKey;index"`
	Name         string         `json:"name" gorm:"not null"`
	ProductId    int            `json:"product_id" gorm:"not null"`
	Sku          string         `json:"sku" gorm:"not null"`
	TotalPrice   float64        `json:"total_price" gorm:"not null"`
	TotalProduct float64        `json:"total_product" gorm:"not null"`
	Description  string         `json:"description"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
