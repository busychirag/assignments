package models

import (
	"time"

	"gorm.io/gorm"
)

// Account represents a bank account
type Account struct {
	gorm.Model
	BranchID       uint            `json:"branch_id"`
	AccountNumber  string          `json:"account_number" gorm:"unique"`
	Balance        float64         `json:"balance"`
	Type           string          `json:"type"`   // e.g., Savings, Checking
	Status         string          `json:"status"` // e.g., Active, Dormant
	OpenedDate     time.Time       `json:"opened_date"`
	InterestRate   float64         `json:"interest_rate"`
	Currency       string          `json:"currency"`
	AccountHolders []AccountHolder `json:"account_holders" gorm:"foreignKey:AccountID"`
	Transactions   []Transaction   `json:"transactions_from" gorm:"foreignKey:FromAccountID"`
	TxRecieved     []Transaction   `json:"transactions_to" gorm:"foreignKey:ToAccountID"`
	Cards          []Card          `json:"cards" gorm:"foreignKey:AccountID"`
	Beneficiaries  []Beneficiary   `json:"beneficiaries" gorm:"foreignKey:AccountID"`
}

// Loan represents a loan account
type Loan struct {
	gorm.Model
	BranchID           uint            `json:"branch_id"`
	PrimaryBorrowerID  uint            `json:"primary_borrower_id"`
	PrincipalAmount    float64         `json:"principal_amount"`
	InterestRate       float64         `json:"interest_rate"`
	EMIAmount          float64         `json:"emi_amount"`
	OutstandingBalance float64         `json:"outstanding_balance"`
	StartDate          time.Time       `json:"start_date"`
	EndDate            time.Time       `json:"end_date"`
	Status             string          `json:"status"` // e.g., Active, Closed, Defaulted
	Collaterals        []Collateral    `json:"collaterals" gorm:"foreignKey:LoanID"`
	Payments           []LoanPayment   `json:"payments" gorm:"foreignKey:LoanID"`
	Guarantors         []LoanGuarantor `json:"guarantors" gorm:"foreignKey:LoanID"`
}

// LoanPayment represents a payment made towards a loan
type LoanPayment struct {
	gorm.Model
	LoanID      uint      `json:"loan_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Mode        string    `json:"mode"`
	Status      string    `json:"status"`
}

// Collateral represents an asset secured against a loan
type Collateral struct {
	gorm.Model
	LoanID uint    `json:"loan_id"`
	Type   string  `json:"type"` // e.g., RealEstate, Vehicle
	Value  float64 `json:"value"`
}

// LoanGuarantor represents a guarantor for a loan
type LoanGuarantor struct {
	gorm.Model
	LoanID          uint     `json:"loan_id"`
	CustomerID      uint     `json:"customer_id"`
	LiabilityAmount float64  `json:"liability_amount"`
	Customer        Customer `json:"customer" gorm:"foreignKey:CustomerID"` // Link back to customer info if needed
}
