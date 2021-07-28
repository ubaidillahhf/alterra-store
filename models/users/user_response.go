package users

import (
	"alterra_store/models/base"
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Token     string         `json:"token"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserResponseTest struct {
	base.BaseResponse
	Data []User `json:"data"`
}
