package v1

import (
	"github.com/dalmarcogd/digital-account/transactions/brokers/events"
	"github.com/dalmarcogd/digital-account/transactions/brokers/rabbit"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

type TransactionCreateRequest struct {
	AccountId       string  `json:"account_id" validate:"required"`
	OperationTypeId int     `json:"operation_type_id" validate:"oneof=1 2 3 4"`
	Amount          float64 `json:"amount" validate:"required"`
}

type TransactionCreateResponse struct {
	TransactionId   string  `json:"transaction_id" validate:"required"`
	AccountId       string  `json:"account_id" validate:"required"`
	OperationTypeId int     `json:"operation_type_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
}

func TransactionsCreateV1Handler(c echo.Context) error {
	transactionRequest := new(TransactionCreateRequest)
	if err := c.Bind(transactionRequest); err != nil {
		return err
	}
	if err := c.Validate(transactionRequest); err != nil {
		return err
	}

	uid, _ := uuid.NewUUID()
	event := events.NewTransactionCreateEvent(uid.String(), transactionRequest.AccountId, transactionRequest.OperationTypeId, transactionRequest.Amount)

	if err := rabbit.NewRabbit().Publish(event); err != nil {
		return err
	}

	transactionResponse := new(TransactionCreateResponse)
	transactionResponse.TransactionId = event.TransactionId
	transactionResponse.AccountId = event.AccountId
	transactionResponse.OperationTypeId = event.OperationTypeId
	transactionResponse.Amount = event.Amount
	return c.JSON(http.StatusOK, transactionResponse)
}
