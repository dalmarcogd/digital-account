package v1

import (
	"fmt"
	"github.com/dalmarcogd/digital-account/accounts/brokers/events"
	"github.com/dalmarcogd/digital-account/accounts/brokers/rabbit"
	"github.com/dalmarcogd/digital-account/accounts/cache"
	"github.com/dalmarcogd/digital-account/accounts/database"
	"github.com/dalmarcogd/digital-account/accounts/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

const DefaultAvailableCreditLimit = 5000


type AccountsCreateRequest struct {
	DocumentNumber       string  `json:"document_number" validate:"required"`
	AvailableCreditLimit float64 `json:"available_credit_limit"`
}

type AccountsCreateResponse struct {
	AccountId            string  `json:"account_id" validate:"required"`
	DocumentNumber       string  `json:"document_number" validate:"required"`
	AvailableCreditLimit float64 `json:"available_credit_limit" validate:"required"`
}

func AccountsCreateV1Handler(c echo.Context) error {
	accountRequest := new(AccountsCreateRequest)
	if err := c.Bind(accountRequest); err != nil {
		return err
	}

	if err := c.Validate(accountRequest); err != nil {
		return err
	}

	uid, _ := uuid.NewUUID()
	if accountRequest.AvailableCreditLimit == 0 {
		accountRequest.AvailableCreditLimit = DefaultAvailableCreditLimit
	}
	event := events.NewAccountCreateEvent(uid.String(), accountRequest.DocumentNumber, accountRequest.AvailableCreditLimit)

	data, err := utils.NewJsonConverter().Encode(event)
	if err != nil {
		return err
	}
	if err := cache.GetConnection().Set(fmt.Sprintf("account-%s", event.AccountId), string(data), 900, 0, false, false); err != nil {
		return err
	}

	if err := rabbit.NewRabbit().Publish(event); err != nil {
		return err
	}

	accountResponse := new(AccountsCreateResponse)
	accountResponse.AccountId = event.AccountId
	accountResponse.DocumentNumber = event.DocumentNumber
	accountResponse.AvailableCreditLimit = event.AvailableCreditLimit
	return c.JSON(http.StatusOK, accountResponse)
}

func AccountsGetV1Handler(c echo.Context) error {
	accountId := c.Param("accountId")

	data, err := cache.GetConnection().Get(fmt.Sprintf("account-%s", accountId))
	if err != nil {
		return err
	}

	event := events.NewAccountCreateEvent("", "", 0)

	if data != nil {
		if err := utils.NewJsonConverter().Decode(data, event); err != nil {
			return err
		}
	} else {
		account, err := database.GetAccountById(accountId)
		if err != nil {
			return err
		}
		event.AccountId = account.Id
		event.DocumentNumber = account.DocumentNumber
		event.AvailableCreditLimit = account.AvailableCreditLimit
	}

	accountResponse := new(AccountsCreateResponse)
	accountResponse.AccountId = event.AccountId
	accountResponse.DocumentNumber = event.DocumentNumber
	accountResponse.AvailableCreditLimit = event.AvailableCreditLimit
	return c.JSON(http.StatusOK, accountResponse)
}
