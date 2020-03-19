package database

import (
	"math"
	"time"
)

func CreateTransaction(transactionId string, accountId string, operationTypeId int, amount float64) error {
	transactionModel := TransactionModel{}
	transactionModel.Id = transactionId
	transactionModel.AccountId = accountId
	transactionModel.OperationType = operationTypeId
	if operationTypeId == CompraAVista || operationTypeId == CompraParcelada || operationTypeId == Saque {
		if !math.Signbit(amount) {
			amount = amount * -1
		}
	} else if operationTypeId == Pagamento {
		if math.Signbit(amount) {
			amount = math.Abs(amount)
		}
	}
	transactionModel.Amount =amount
	transactionModel.EventDate = time.Now().UTC()
	return GetConnection().Save(&transactionModel).Error
}
