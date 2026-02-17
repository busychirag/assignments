package models

type Loan struct {
	BaseModel
	CustomerID      uint64   `json:"customer_id" gorm:"not null"`
	BranchID        uint64   `json:"branch_id" gorm:"not null"`
	Amount          float64  `json:"amount" gorm:"not null"`
	InterestRate    float64  `json:"interest_rate" gorm:"default:12"`
	RemainingAmount float64  `json:"remaining_amount"`
	Customer        Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	Branch          Branch   `json:"branch,omitempty" gorm:"foreignKey:BranchID"`
}
