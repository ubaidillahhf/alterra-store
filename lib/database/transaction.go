package database

import (
	"alterra_store/configs"
	"alterra_store/models/transactions"
)

func CreateTransaction(transactionCreate transactions.TransactionStruct) (transactions.Transaction, error) {
	var transactionsDB transactions.Transaction

	transactionsDB.Name = transactionCreate.Name
	transactionsDB.Status = transactionCreate.Status
	transactionsDB.TransactionId = transactionCreate.TransactionId
	transactionsDB.Description = transactionCreate.Description

	err := configs.DB.Create(&transactionsDB).Error
	if err != nil {
		return transactionsDB, err
	}
	return transactionsDB, nil
}

func GetTransactionDetail(transactionId int) (transactions.Transaction, error) {
	var transactionsDB transactions.Transaction
	err := configs.DB.First(&transactionsDB, transactionId).Error

	if err != nil {
		return transactionsDB, err
	}
	return transactionsDB, nil
}
