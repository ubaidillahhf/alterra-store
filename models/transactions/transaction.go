package transactions

type TransactionStruct struct {
	TransactionId int    `json:"transaction_id" gorm:"index;not null"`
	Name          string `json:"name" gorm:"not null"`
	Status        string `json:"status" gorm:"not null"`
	Description   string `json:"description"`
}
