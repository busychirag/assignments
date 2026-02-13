package services

import (
	"errors"
	"fmt"

	"github.com/busychirag/assignments/tree/main/Assignment2/initializers"
	"github.com/busychirag/assignments/tree/main/Assignment2/models"
	"gorm.io/gorm"
)

type LoanDetails struct {
	models.Loan
	LoanPending      float64 `json:"loan_pending"`
	InterestThisYear float64 `json:"interest_this_year"`
	TotalPayable     float64 `json:"total_payable"`
}

func GetAllLoans() ([]models.Loan, error) {
	var loans []models.Loan
	result := initializers.DB.Preload("Customer").Preload("Branch").Find(&loans)
	return loans, result.Error
}

func GetLoanByID(id uint64) (LoanDetails, error) {
	var loan models.Loan
	if err := initializers.DB.Preload("Customer").Preload("Branch").First(&loan, id).Error; err != nil {
		return LoanDetails{}, err
	}

	interestThisYear := loan.RemainingAmount * (loan.InterestRate / 100)
	totalPayable := loan.RemainingAmount + interestThisYear

	details := LoanDetails{
		Loan:             loan,
		LoanPending:      loan.RemainingAmount,
		InterestThisYear: interestThisYear,
		TotalPayable:     totalPayable,
	}

	return details, nil
}

func GetLoansByCustomerID(customerID uint64) ([]models.Loan, error) {
	var loans []models.Loan
	result := initializers.DB.Preload("Branch").Where("customer_id = ?", customerID).Find(&loans)
	return loans, result.Error
}

func CreateLoan(loan *models.Loan) error {
	if loan.Amount <= 0 {
		return errors.New("loan amount must be greater than zero")
	}

	loan.InterestRate = 12
	loan.RemainingAmount = loan.Amount

	result := initializers.DB.Create(loan)
	return result.Error
}

func RepayLoan(loanID uint64, amount float64) (models.Loan, error) {
	var loan models.Loan
	if err := initializers.DB.First(&loan, loanID).Error; err != nil {
		return loan, err
	}

	if amount <= 0 {
		return loan, errors.New("repayment amount must be greater than zero")
	}

	if loan.RemainingAmount <= 0 {
		return loan, errors.New("loan is already fully repaid")
	}

	if amount > loan.RemainingAmount {
		amount = loan.RemainingAmount // cap at remaining
	}

	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		loan.RemainingAmount -= amount
		if err := tx.Save(&loan).Error; err != nil {
			return err
		}

		// Record as transaction on the customer's first account (if exists)
		var account models.Account
		if err := tx.Where("customer_id = ?", loan.CustomerID).First(&account).Error; err == nil {
			transaction := models.Transaction{
				AccountID:   account.ID,
				Type:        "loan_repayment",
				Amount:      amount,
				Description: fmt.Sprintf("Loan #%d repayment of %.2f", loanID, amount),
			}
			if err := tx.Create(&transaction).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return loan, err
}
