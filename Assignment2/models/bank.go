package models

type Bank struct {
	BaseModel
	Name     string   `json:"name" gorm:"not null"`
	Branches []Branch `json:"branches,omitempty" gorm:"foreignKey:BankID"`
}
