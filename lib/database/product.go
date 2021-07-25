package database

import (
	"alterra_store/configs"
	"alterra_store/models/products"
)

func CreateProduct(productCreate products.ProductStruct) (products.Product, error) {
	var productDB products.Product

	productDB.Name = productCreate.Name
	productDB.CategoryId = productCreate.CategoryId
	productDB.Sku = productCreate.Sku
	productDB.Price = productCreate.Price
	productDB.Description = productCreate.Description

	err := configs.DB.Create(&productDB).Error
	if err != nil {
		return productDB, err
	}
	return productDB, nil
}
