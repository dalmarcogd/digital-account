package events

type accountCreateEvent struct {
	*EventBase
	AccountId            string  `json:"account_id" validate:"required"`
	DocumentNumber       string  `json:"document_number" validate:"required"`
	AvailableCreditLimit float64 `json:"available_credit_limit" validate:"required"`
}

func (a accountCreateEvent) GetChannel() string {
	return a.Name
}

func NewAccountCreateEvent(id string, documentNumber string, availableCreditLimit float64) *accountCreateEvent {
	return &accountCreateEvent{EventBase: NewEventBase("AccountCreateEvent"), AccountId: id, DocumentNumber: documentNumber, AvailableCreditLimit: availableCreditLimit}
}
