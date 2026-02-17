package services

import (
	"github.com/busychirag/assignments/tree/main/Assignment2/initializers"
	"github.com/busychirag/assignments/tree/main/Assignment2/models"
	"gorm.io/gorm"
)

func GetAllBanks() ([]models.Bank, error) {
	var banks []models.Bank
	result := initializers.DB.Find(&banks)
	return banks, result.Error
}

func GetBankByID(id uint64) (models.Bank, error) {
	var bank models.Bank
	result := initializers.DB.Preload("Branches").First(&bank, id)
	return bank, result.Error
}

func CreateBank(bank *models.Bank) error {
	result := initializers.DB.Create(bank)
	return result.Error
}

func UpdateBank(bank *models.Bank) error {
	result := initializers.DB.Save(bank)
	return result.Error
}

func DeleteBank(id uint64) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		// Find all branches of this bank
		var branches []models.Branch
		tx.Where("bank_id = ?", id).Find(&branches)

		// Cascade delete each branch's accounts and loans
		for _, branch := range branches {
			if err := deleteBranchCascade(tx, branch.ID); err != nil {
				return err
			}
		}

		// Delete the bank itself
		return tx.Delete(&models.Bank{}, id).Error
	})
}

// deleteBranchCascade deletes a branch and all its associated accounts, loans,
// transactions, and joint account holders within the given transaction.
func deleteBranchCascade(tx *gorm.DB, branchID uint64) error {
	// Find all accounts of this branch
	var accounts []models.Account
	tx.Where("branch_id = ?", branchID).Find(&accounts)

	for _, account := range accounts {
		// Delete transactions of this account
		if err := tx.Where("account_id = ?", account.ID).Delete(&models.Transaction{}).Error; err != nil {
			return err
		}
		// Delete joint account holders of this account
		if err := tx.Where("account_id = ?", account.ID).Delete(&models.JointAccountHolder{}).Error; err != nil {
			return err
		}
	}

	// Delete all accounts of this branch
	if err := tx.Where("branch_id = ?", branchID).Delete(&models.Account{}).Error; err != nil {
		return err
	}

	// Delete all loans of this branch
	if err := tx.Where("branch_id = ?", branchID).Delete(&models.Loan{}).Error; err != nil {
		return err
	}

	// Delete the branch itself
	return tx.Delete(&models.Branch{}, branchID).Error
}
