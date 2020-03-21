package events

import "math"

const (
	CompraAVista    = 1
	CompraParcelada = 2
	Saque           = 3
	Pagamento       = 4
)

type transactionCreateEvent struct {
	*EventBase
	TransactionId   string  `json:"transaction_id" validate:"required"`
	AccountId       string  `json:"account_id" validate:"required"`
	OperationTypeId int     `json:"operation_type_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
}

func (a transactionCreateEvent) GetChannel() string {
	return a.Name
}

func NewTransactionCreateEvent(id string, accountId string, operationTypeId int, amount float64) *transactionCreateEvent {
	if operationTypeId == CompraAVista || operationTypeId == CompraParcelada || operationTypeId == Saque {
		if !math.Signbit(amount) {
			amount = amount * -1
		}
	} else if operationTypeId == Pagamento {
		if math.Signbit(amount) {
			amount = math.Abs(amount)
		}
	}
	return &transactionCreateEvent{EventBase: NewEventBase("TransactionCreateEvent"), TransactionId: id, AccountId: accountId, OperationTypeId: operationTypeId, Amount: amount}
}
