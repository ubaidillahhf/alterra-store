package carts

type CartStruct struct {
	Name         string  `json:"name" gorm:"not null"`
	ProductId    int     `json:"product_id" gorm:"index;not null"`
	Sku          string  `json:"sku" gorm:"not null"`
	TotalPrice   float64 `json:"total_price" gorm:"not null"`
	TotalProduct float64 `json:"total_product" gorm:"not null"`
	Description  string  `json:"description"`
}
