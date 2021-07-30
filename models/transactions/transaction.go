package transactions

import (
	"time"
)

type TransactionStruct struct {
	Number            string                    `json:"number" gorm:"not null"`
	UserId            int                       `json:"user_id" gorm:"not null"`
	Date              time.Time                 `json:"date" gorm:"not null"`
	Status            string                    `json:"status" gorm:"not null"`
	TotalPrice        float64                   `json:"total_price" gorm:"not null"`
	TransactionDetail []TransactionDetailStruct `json:"transaction_detail" gorm:"not null"`
}
