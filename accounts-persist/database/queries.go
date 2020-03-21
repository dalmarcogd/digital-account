package database

func CreateAccount(accountId string, documentNumber string, availableCreditLimit float64) error {
	account := AccountModel{}
	account.Id = accountId
	account.DocumentNumber = documentNumber
	account.AvailableCreditLimit = availableCreditLimit
	return GetConnection().Save(&account).Error
}
