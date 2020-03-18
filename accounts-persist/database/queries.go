package database

func CreateAccount(accountId string, documentNumber string) error {
	account := AccountModel{}
	account.Id = accountId
	account.DocumentNumber = documentNumber
	return GetConnection().Save(&account).Error
}
