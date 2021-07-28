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

func GetProductDetail(categoryId int) (products.Product, error) {
	var productDB products.Product
	err := configs.DB.First(&productDB, categoryId).Error

	if err != nil {
		return productDB, err
	}
	return productDB, nil
}

func EditProduct(productEdit products.ProductStruct, categoryId int) (products.Product, error) {
	var productDB products.Product
	err := configs.DB.First(&productDB, categoryId).Error

	productDB.Name = productEdit.Name
	productDB.Description = productEdit.Description

	configs.DB.Save(productDB)

	if err != nil {
		return productDB, err
	}

	return productDB, nil
}

func DeleteProduct(categoryId int) (products.Product, error) {
	var productDB products.Product
	err := configs.DB.Where("id = ?", categoryId).Delete(&productDB).Error

	if err != nil {
		return productDB, err
	}
	return productDB, nil
}
