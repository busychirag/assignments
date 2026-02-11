package models

import (
	"time"

	"gorm.io/gorm"
)

// Transaction represents a financial transaction between accounts
type Transaction struct {
	gorm.Model
	FromAccountID   uint      `json:"from_account_id"`
	ToAccountID     uint      `json:"to_account_id"`
	Amount          float64   `json:"amount"`
	Type            string    `json:"type"` // e.g., Transfer, Withdrawal, Deposit
	Mode            string    `json:"mode"` // e.g., Online, ATM
	Status          string    `json:"status"`
	TransactionDate time.Time `json:"transaction_date"`
	ReferenceNumber string    `json:"reference_number" gorm:"unique"`
	FromAccount     *Account  `json:"from_account" gorm:"foreignKey:FromAccountID"`
	ToAccount       *Account  `json:"to_account" gorm:"foreignKey:ToAccountID"`
}

// Card represents a debit or credit card linked to an account
type Card struct {
	gorm.Model
	AccountID  uint      `json:"account_id"`
	CardNumber string    `json:"card_number" gorm:"unique"`
	Type       string    `json:"type"` // e.g., Debit, Credit
	Expiry     time.Time `json:"expiry"`
	IssuedDate time.Time `json:"issued_date"`
	Status     string    `json:"status"`
}
