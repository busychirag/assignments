package models

type Customer struct {
	BaseModel
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"uniqueIndex;not null"`
	Phone string `json:"phone"`
	Loans []Loan `json:"loans,omitempty" gorm:"foreignKey:CustomerID"`
}
