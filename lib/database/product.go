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

func GetProductAll() (dataResult []products.Product, err error) {
	err = configs.DB.Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetProductDetail(productId int) (products.Product, error) {
	var productDB products.Product
	err := configs.DB.First(&productDB, productId).Error

	if err != nil {
		return productDB, err
	}
	return productDB, nil
}

func EditProduct(productEdit products.ProductStruct, productId int) (products.Product, error) {
	var productDB products.Product
	err := configs.DB.First(&productDB, productId).Error

	productDB.Name = productEdit.Name
	productDB.CategoryId = productEdit.CategoryId
	productDB.Sku = productEdit.Sku
	productDB.Price = productEdit.Price
	productDB.Description = productEdit.Description

	configs.DB.Save(productDB)

	if err != nil {
		return productDB, err
	}

	return productDB, nil
}

func DeleteProduct(productId int) (products.Product, error) {
	var productDB products.Product
	err := configs.DB.Where("id = ?", productId).Delete(&productDB).Error

	if err != nil {
		return productDB, err
	}
	return productDB, nil
}
