package services

import (
	"fmt"
	"github.com/dalmarcogd/digital-account/transactions/database"
	"github.com/dalmarcogd/digital-account/transactions/environments"
	"github.com/dalmarcogd/digital-account/transactions/errors"
	"github.com/dalmarcogd/digital-account/transactions/utils"
	"io/ioutil"
	"net/http"
)

func GetAvailableCreditLimitByAccount(accountId string) (float64, error) {
	env := environments.GetEnvironment()
	response, err := http.Get(fmt.Sprintf("%s/accounts-api/v1/accounts/%s", env.AccountsUtl, accountId))
	if err != nil {
		return 0, err
	}

	if response.StatusCode != http.StatusOK {
		return 0, errors.NewError(http.StatusBadRequest, "Account not found.", nil)
	}

	body := make(map[string]interface{})
	bodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	if err := utils.NewJsonConverter().Decode(bodyData, &body); err != nil {
		return 0, err
	}

	if _, ok := body["available_credit_limit"]; !ok {
		return 0, errors.NewError(http.StatusBadRequest, "Unable to find available_credit_limit", nil)
	}

	availableCreditLimitData := body["available_credit_limit"]

	availableCreditLimit := 0.0
	if  availableCreditLimitData != nil {
		availableCreditLimit = availableCreditLimitData.(float64)
	}

	transactions, err := database.GetTransactionsByAccountId(accountId)
	if err != nil {
		return 0, err
	}
	current := 0.0

	for _, transaction := range transactions {
		current += transaction.Amount
	}

	return availableCreditLimit + current, nil

}
