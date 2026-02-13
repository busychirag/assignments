package models

type Branch struct {
	BaseModel
	BankID   uint64    `json:"bank_id" gorm:"not null"`
	Name     string    `json:"name" gorm:"not null"`
	Address  string    `json:"address"`
	Bank     Bank      `json:"bank,omitempty" gorm:"foreignKey:BankID"`
	Accounts []Account `json:"accounts,omitempty" gorm:"foreignKey:BranchID"`
	Loans    []Loan    `json:"loans,omitempty" gorm:"foreignKey:BranchID"`
}
