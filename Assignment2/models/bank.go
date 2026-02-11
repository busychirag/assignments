package models

import "gorm.io/gorm"

// Bank represents the bank organization
type Bank struct {
	gorm.Model
	Name        string   `json:"name"`
	SwiftCode   string   `json:"swift_code" gorm:"unique"`
	TotalAssets float64  `json:"total_assets"`
	Branches    []Branch `json:"branches" gorm:"foreignKey:BankID"`
}

// Branch represents a bank branch
type Branch struct {
	gorm.Model
	BankID     uint       `json:"bank_id"`
	Name       string     `json:"name"`
	BranchCode string     `json:"branch_code" gorm:"unique"`
	City       string     `json:"city"`
	Employees  []Employee `json:"employees" gorm:"foreignKey:BranchID"`
	Accounts   []Account  `json:"accounts" gorm:"foreignKey:BranchID"`
	Loans      []Loan     `json:"loans" gorm:"foreignKey:BranchID"`
}

// Employee represents a bank employee
type Employee struct {
	gorm.Model
	BranchID     uint       `json:"branch_id"`
	SupervisorID *uint      `json:"supervisor_id"` // Self-referential
	Name         string     `json:"name"`
	Designation  string     `json:"designation"`
	Supervisor   *Employee  `json:"supervisor" gorm:"foreignKey:SupervisorID"`
	Subordinates []Employee `json:"subordinates" gorm:"foreignKey:SupervisorID"`
}
