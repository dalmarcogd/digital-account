package database

import "time"

const (
	CompraAVista    = 1
	CompraParcelada = 2
	Saque           = 3
	Pagamento       = 4
)

type TransactionModel struct {
	Id            string `gorm:"primary_key"`
	AccountId     string
	OperationType int
	Amount        float64
	EventDate     time.Time
}

func (TransactionModel) TableName() string {
	return "transactions"
}
