package database

import (
	"alterra_store/configs"
	"alterra_store/models/carts"
)

func CreateCart(createCart carts.CartStruct) (carts.Cart, error) {
	var cartDb carts.Cart

	cartDb.Name = createCart.Name
	cartDb.Description = createCart.Description
	cartDb.TotalPrice = createCart.TotalPrice
	cartDb.TotalProduct = createCart.TotalProduct
	cartDb.Sku = createCart.Sku
	cartDb.ProductId = createCart.ProductId

	err := configs.DB.Create(&cartDb).Error
	if err != nil {
		return cartDb, err
	}
	return cartDb, nil
}

func DeleteCart(cartId int) (carts.Cart, error) {
	var cartDb carts.Cart
	err := configs.DB.Where("id = ?", cartId).Delete(&cartDb).Error

	if err != nil {
		return cartDb, err
	}
	return cartDb, nil
}

func GetCartAll() (dataResult []carts.Cart, err error) {
	err = configs.DB.Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}
