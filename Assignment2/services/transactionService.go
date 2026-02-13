package services

import (
	"github.com/busychirag/assignments/tree/main/Assignment2/initializers"
	"github.com/busychirag/assignments/tree/main/Assignment2/models"
)

func GetTransactionsByAccountID(accountID uint64) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := initializers.DB.Where("account_id = ?", accountID).Order("created_at desc").Find(&transactions)
	return transactions, result.Error
}
