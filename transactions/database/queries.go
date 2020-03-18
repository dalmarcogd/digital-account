package database

import (
	"github.com/dalmarcogd/digital-account/transactions/errors"
	"net/http"
)

func GetAccountById(accountId string) (TransactionModel, error) {
	account := TransactionModel{}
	GetConnection().Where("id = ?", accountId).First(&account)
	if account.Id != "" {
		return account, nil
	}
	return account, errors.NewError(http.StatusNotFound, "Account not found", nil)
}