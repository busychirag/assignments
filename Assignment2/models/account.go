package models

type Account struct {
	BaseModel
	BranchID     uint64               `json:"branch_id" gorm:"not null"`
	AccountType  string               `json:"account_type" gorm:"default:'savings'"`
	Balance      float64              `json:"balance" gorm:"default:0"`
	Branch       Branch               `json:"branch,omitempty" gorm:"foreignKey:BranchID"`
	Transactions []Transaction        `json:"transactions,omitempty" gorm:"foreignKey:AccountID"`
	Holders      []JointAccountHolder `json:"holders,omitempty" gorm:"foreignKey:AccountID"`
}
