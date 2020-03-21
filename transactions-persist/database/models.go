package database

import "time"

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
