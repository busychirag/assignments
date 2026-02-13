package services

import (
	"errors"
	"fmt"

	"github.com/busychirag/assignments/tree/main/Assignment2/initializers"
	"github.com/busychirag/assignments/tree/main/Assignment2/models"
	"gorm.io/gorm"
)

func GetAllAccounts() ([]models.Account, error) {
	var accounts []models.Account
	result := initializers.DB.Preload("Branch").Preload("Holders").Preload("Holders.Customer").Find(&accounts)
	return accounts, result.Error
}

func GetAccountByID(id uint64) (models.Account, error) {
	var account models.Account
	result := initializers.DB.Preload("Branch").Preload("Holders").Preload("Holders.Customer").First(&account, id)
	return account, result.Error
}

func GetAccountsByCustomerID(customerID uint64) ([]models.Account, error) {
	// Find all accounts where this customer is a holder (primary or joint)
	var holdings []models.JointAccountHolder
	initializers.DB.Where("customer_id = ?", customerID).Find(&holdings)

	var accounts []models.Account
	for _, h := range holdings {
		var acc models.Account
		if err := initializers.DB.Preload("Branch").Preload("Holders").Preload("Holders.Customer").First(&acc, h.AccountID).Error; err == nil {
			accounts = append(accounts, acc)
		}
	}

	return accounts, nil
}

func CreateAccount(account *models.Account, customerID uint64) error {
	// Verify customer exists
	var customer models.Customer
	if err := initializers.DB.First(&customer, customerID).Error; err != nil {
		return errors.New("customer not found")
	}

	account.AccountType = "savings"
	account.Balance = 0
	if err := initializers.DB.Create(account).Error; err != nil {
		return err
	}

	// Register primary holder in join table
	holder := models.JointAccountHolder{
		AccountID:  account.ID,
		CustomerID: customerID,
		IsPrimary:  true,
	}
	return initializers.DB.Create(&holder).Error
}

func AddJointHolder(accountID uint64, customerID uint64) (models.JointAccountHolder, error) {
	// Verify account exists
	var account models.Account
	if err := initializers.DB.First(&account, accountID).Error; err != nil {
		return models.JointAccountHolder{}, errors.New("account not found")
	}

	// Verify customer exists
	var customer models.Customer
	if err := initializers.DB.First(&customer, customerID).Error; err != nil {
		return models.JointAccountHolder{}, errors.New("customer not found")
	}

	// Check if already a holder
	var existing models.JointAccountHolder
	if err := initializers.DB.Where("account_id = ? AND customer_id = ?", accountID, customerID).First(&existing).Error; err == nil {
		return existing, errors.New("customer is already a holder of this account")
	}

	holder := models.JointAccountHolder{
		AccountID:  accountID,
		CustomerID: customerID,
		IsPrimary:  false,
	}
	if err := initializers.DB.Create(&holder).Error; err != nil {
		return holder, err
	}

	initializers.DB.Preload("Customer").First(&holder, holder.ID)
	return holder, nil
}

func RemoveJointHolder(accountID uint64, customerID uint64) error {
	var holder models.JointAccountHolder
	if err := initializers.DB.Where("account_id = ? AND customer_id = ?", accountID, customerID).First(&holder).Error; err != nil {
		return errors.New("holder not found")
	}

	if holder.IsPrimary {
		return errors.New("cannot remove primary account holder")
	}

	return initializers.DB.Delete(&holder).Error
}

func GetAccountHolders(accountID uint64) ([]models.JointAccountHolder, error) {
	var holders []models.JointAccountHolder
	result := initializers.DB.Preload("Customer").Where("account_id = ?", accountID).Find(&holders)
	return holders, result.Error
}

func Deposit(accountID uint64, amount float64) (models.Account, error) {
	var account models.Account
	if err := initializers.DB.First(&account, accountID).Error; err != nil {
		return account, err
	}

	if amount <= 0 {
		return account, errors.New("deposit amount must be greater than zero")
	}

	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		account.Balance += amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		transaction := models.Transaction{
			AccountID:   accountID,
			Type:        "deposit",
			Amount:      amount,
			Description: fmt.Sprintf("Deposited %.2f", amount),
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil
	})

	return account, err
}

func Withdraw(accountID uint64, amount float64) (models.Account, error) {
	var account models.Account
	if err := initializers.DB.First(&account, accountID).Error; err != nil {
		return account, err
	}

	if amount <= 0 {
		return account, errors.New("withdrawal amount must be greater than zero")
	}

	if account.Balance < amount {
		return account, errors.New("insufficient balance")
	}

	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		account.Balance -= amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		transaction := models.Transaction{
			AccountID:   accountID,
			Type:        "withdrawal",
			Amount:      amount,
			Description: fmt.Sprintf("Withdrew %.2f", amount),
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil
	})

	return account, err
}
