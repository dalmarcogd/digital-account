package database

import (
	"time"
)

func CreateTransaction(transactionId string, accountId string, operationTypeId int, amount float64) error {
	transactionModel := TransactionModel{}
	transactionModel.Id = transactionId
	transactionModel.AccountId = accountId
	transactionModel.OperationType = operationTypeId
	transactionModel.Amount =amount
	transactionModel.EventDate = time.Now().UTC()
	return GetConnection().Save(&transactionModel).Error
}
