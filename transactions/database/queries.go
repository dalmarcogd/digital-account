package database


func GetTransactionsByAccountId(accountId string) ([]TransactionModel, error) {
	var transactions []TransactionModel
	if err := GetConnection().Find(&transactions, TransactionModel{AccountId:accountId}).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}