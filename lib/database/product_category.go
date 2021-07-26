package database

import (
	"alterra_store/configs"
	"alterra_store/models/productCategories"
)

func CreateCategory(categoryCreate productCategories.ProductCategoryStruct) (productCategories.ProductCategory, error) {
	var productCategoryDB productCategories.ProductCategory

	productCategoryDB.Name = categoryCreate.Name
	productCategoryDB.Description = categoryCreate.Description

	err := configs.DB.Create(&productCategoryDB).Error
	if err != nil {
		return productCategoryDB, err
	}
	return productCategoryDB, nil
}

func GetCategoryAll() (dataResult []productCategories.ProductCategory, err error) {
	err = configs.DB.Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetCategoryDetail(categoryId int) (productCategories.ProductCategory, error) {
	var categoryDB productCategories.ProductCategory
	err := configs.DB.First(&categoryDB, categoryId).Error

	if err != nil {
		return categoryDB, err
	}
	return categoryDB, nil
}

func EditCategory(categoryEdit productCategories.ProductCategoryStruct, categoryId int) (productCategories.ProductCategory, error) {
	var categoryDB productCategories.ProductCategory
	err := configs.DB.First(&categoryDB, categoryId).Error

	categoryDB.Name = categoryEdit.Name
	categoryDB.Description = categoryEdit.Description

	configs.DB.Save(&categoryDB)

	if err != nil {
		return categoryDB, err
	}
	return categoryDB, nil
}

func DeleteCategory(categoryId int) (productCategories.ProductCategory, error) {
	var categoryDB productCategories.ProductCategory
	err := configs.DB.Where("id = ?", categoryId).Delete(&categoryDB).Error

	if err != nil {
		return categoryDB, err
	}
	return categoryDB, nil
}
