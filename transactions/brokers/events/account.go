package events

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
	return &transactionCreateEvent{EventBase: NewEventBase("TransactionCreateEvent"), TransactionId: id, AccountId: accountId, OperationTypeId: operationTypeId, Amount: amount}
}
