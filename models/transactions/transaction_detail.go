package transactions

type TransactionDetailStruct struct {
	TransactionId int     `json:"transaction_id" gorm:"index;not null"`
	Name          string  `json:"name" gorm:"not null"`
	CategoryId    int     `json:"category_id" gorm:"not null"`
	Sku           string  `json:"sku" gorm:"not null"`
	Price         float64 `json:"price" gorm:"not null"`
	Description   string  `json:"description"`
}
