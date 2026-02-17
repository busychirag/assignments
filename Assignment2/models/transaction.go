package models

type Transaction struct {
	BaseModel
	AccountID   uint64  `json:"account_id" gorm:"not null"`
	Type        string  `json:"type" gorm:"not null"` // deposit, withdrawal, loan_repayment
	Amount      float64 `json:"amount" gorm:"not null"`
	Description string  `json:"description"`
	Account     Account `json:"account,omitempty" gorm:"foreignKey:AccountID"`
}
