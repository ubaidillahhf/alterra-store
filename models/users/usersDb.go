package users

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
