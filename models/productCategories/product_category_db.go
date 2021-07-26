package productCategories

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	Id          int            `json:"id" gorm:"primaryKey;index"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
