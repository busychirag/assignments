package main

import (
	"github.com/busychirag/assignments/tree/main/Assignment2/initializers"
	"github.com/busychirag/assignments/tree/main/Assignment2/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.Bank{},
		&models.Branch{},
		&models.Employee{},
		&models.Customer{},
		&models.KYCDocument{},
		&models.AccountHolder{},
		&models.Beneficiary{},
		&models.Account{},
		&models.Loan{},
		&models.LoanPayment{},
		&models.Collateral{},
		&models.LoanGuarantor{},
		&models.Transaction{},
		&models.Card{},
	)
}
