package database

import (
	"alterra_store/configs"
	"alterra_store/models/transactions"
)

func CreateTransactionDetail(transactionDetailCreate []transactions.TransactionDetail) ([]transactions.TransactionDetail, error) {

	// Transaction Detail
	// var transactionDetailDB []transactions.TransactionDetail

	// for i := 0; i < len(transactionDetailCreate); i++ {
	// 	transactionDetailDB[i].TransactionId = transactionId
	// 	transactionDetailDB[i].Name = transactionDetailCreate[i].Name
	// 	transactionDetailDB[i].CategoryId = transactionDetailCreate[i].CategoryId
	// 	transactionDetailDB[i].Sku = transactionDetailCreate[i].Sku
	// 	transactionDetailDB[i].Price = transactionDetailCreate[i].Price
	// 	transactionDetailDB[i].Description = transactionDetailCreate[i].Description
	// }

	if err := configs.DB.Create(&transactionDetailCreate).Error; err != nil {
		return transactionDetailCreate, err
	}

	return transactionDetailCreate, nil
}
