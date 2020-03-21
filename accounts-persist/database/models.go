package database

type AccountModel struct {
	Id                   string `gorm:"primary_key"`
	DocumentNumber       string
	AvailableCreditLimit float64
}

func (AccountModel) TableName() string {
	return "accounts"
}
