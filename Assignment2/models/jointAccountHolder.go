package models

type JointAccountHolder struct {
	BaseModel
	AccountID  uint64   `json:"account_id" gorm:"not null;uniqueIndex:idx_account_customer"`
	CustomerID uint64   `json:"customer_id" gorm:"not null;uniqueIndex:idx_account_customer"`
	IsPrimary  bool     `json:"is_primary" gorm:"default:false"`
	Account    Account  `json:"account,omitempty" gorm:"foreignKey:AccountID"`
	Customer   Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
}
