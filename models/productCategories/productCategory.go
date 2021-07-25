package productCategories

type ProductCategoryStruct struct {
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
}
