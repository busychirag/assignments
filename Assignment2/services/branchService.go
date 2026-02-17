package services

import (
	"github.com/busychirag/assignments/tree/main/Assignment2/initializers"
	"github.com/busychirag/assignments/tree/main/Assignment2/models"
	"gorm.io/gorm"
)

func GetAllBranches() ([]models.Branch, error) {
	var branches []models.Branch
	result := initializers.DB.Preload("Bank").Find(&branches)
	return branches, result.Error
}

func GetBranchByID(id uint64) (models.Branch, error) {
	var branch models.Branch
	result := initializers.DB.Preload("Bank").First(&branch, id)
	return branch, result.Error
}

func GetBranchesByBankID(bankID uint64) ([]models.Branch, error) {
	var branches []models.Branch
	result := initializers.DB.Where("bank_id = ?", bankID).Find(&branches)
	return branches, result.Error
}

func CreateBranch(branch *models.Branch) error {
	result := initializers.DB.Create(branch)
	return result.Error
}

func UpdateBranch(branch *models.Branch) error {
	result := initializers.DB.Save(branch)
	return result.Error
}

func DeleteBranch(id uint64) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		return deleteBranchCascade(tx, id)
	})
}
