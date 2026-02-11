package models

import (
	"time"

	"gorm.io/gorm"
)

// Customer represents a bank customer
type Customer struct {
	gorm.Model
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	TaxID          string          `json:"tax_id" gorm:"unique"`
	CreditScore    int             `json:"credit_score"`
	KYCDocuments   []KYCDocument   `json:"kyc_documents" gorm:"foreignKey:CustomerID"`
	AccountHolders []AccountHolder `json:"account_holders" gorm:"foreignKey:CustomerID"`
	Loans          []Loan          `json:"loans" gorm:"foreignKey:PrimaryBorrowerID"`
	LoanGuarantors []LoanGuarantor `json:"loan_guarantors" gorm:"foreignKey:CustomerID"`
}

// KYCDocument represents a customer's KYC document
type KYCDocument struct {
	gorm.Model
	CustomerID         uint      `json:"customer_id"`
	DocType            string    `json:"doc_type"`
	VerificationStatus string    `json:"verification_status"`
	IssuedBy           string    `json:"issued_by"`
	ExpiryDate         time.Time `json:"expiry_date"`
}

// AccountHolder links a customer to an account (Many-to-Many relationship primarily handled via this join model)
type AccountHolder struct {
	gorm.Model
	CustomerID   uint     `json:"customer_id"`
	AccountID    uint     `json:"account_id"`
	Role         string   `json:"role"` // e.g., Primary, Joint
	AccessRights string   `json:"access_rights"`
	Customer     Customer `json:"customer" gorm:"foreignKey:CustomerID"`
	Account      Account  `json:"account" gorm:"foreignKey:AccountID"`
}

// Beneficiary represents a beneficiary for an account
type Beneficiary struct {
	gorm.Model
	AccountID uint   `json:"account_id"`
	Name      string `json:"name"`
	Relation  string `json:"relation"`
}
