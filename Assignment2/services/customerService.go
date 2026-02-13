package services

import (
	"github.com/busychirag/assignments/tree/main/Assignment2/initializers"
	"github.com/busychirag/assignments/tree/main/Assignment2/models"
)

func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	result := initializers.DB.Find(&customers)
	return customers, result.Error
}

func GetCustomerByID(id uint64) (models.Customer, error) {
	var customer models.Customer
	result := initializers.DB.Preload("Loans").First(&customer, id)
	return customer, result.Error
}

func CreateCustomer(customer *models.Customer) error {
	result := initializers.DB.Create(customer)
	return result.Error
}

func UpdateCustomer(customer *models.Customer) error {
	result := initializers.DB.Save(customer)
	return result.Error
}

func DeleteCustomer(id uint64) error {
	result := initializers.DB.Delete(&models.Customer{}, id)
	return result.Error
}
