package products

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          int            `json:"id" gorm:"primaryKey;index"`
	Name        string         `json:"name" gorm:"not null"`
	CategoryId  int            `json:"category_id" gorm:"index;not null"`
	Sku         string         `json:"sku" gorm:"not null"`
	Price       float64        `json:"price" gorm:"not null"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
