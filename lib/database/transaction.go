package database

import (
	"alterra_store/configs"
	"alterra_store/models/transactions"
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

func CreateTransaction(transactionCreate transactions.TransactionStruct) (transactions.Transaction, error) {
	var transactionsDB transactions.Transaction
	currentTime := time.Now()
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()

	// Transaction Count
	transactionAll := []transactions.Transaction{}
	numberTransaction := configs.DB.Where("created_at BETWEEN ? AND ?", now.BeginningOfMonth(), currentTime).Find(&transactionAll)
	countTransactionInMonth := numberTransaction.RowsAffected

	transactionsDB.Number = fmt.Sprintf("TRX-%d%d%d-%04d", year, month, day, countTransactionInMonth)
	transactionsDB.UserId = transactionCreate.UserId
	transactionsDB.Date = currentTime
	transactionsDB.Status = transactionCreate.Status
	transactionsDB.TotalPrice = transactionCreate.TotalPrice

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
