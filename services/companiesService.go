package services

import (
	"errors"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
)

// CreateCompany crea una nueva compañía en la base de datos.
func CreateCompany(company models.Companies) (models.Companies, error) {
	// Verificar que el campo 'Name' no esté vacío
	if company.Name == "" {
		return models.Companies{}, errors.New("the 'Name' field is required")
	}

	result := database.DB.Create(&company)
	return company, result.Error
}

// FetchCompanies recupera compañías con paginación de la base de datos.
func FetchCompanies(page int, pageSize int) ([]models.Companies, int64, error) {
	var companies []models.Companies
	var totalRows int64 = 0
	offset := (page - 1) * pageSize
	result := database.DB.Offset(offset).Limit(pageSize).Find(&companies)
	database.DB.Model(&models.Companies{}).Count(&totalRows)
	return companies, totalRows, result.Error
}

// GetCompanyByID busca una compañía por su ID.
func GetCompanyByID(id uint) (models.Companies, error) {
	var company models.Companies
	result := database.DB.First(&company, id)
	return company, result.Error
}

func UpdateCompany(company models.Companies, id uint) (models.Companies, error) {
	// Primero, verifica si existe la compañía que se intenta actualizar.
	var existing models.Companies
	if err := database.DB.First(&existing, id).Error; err != nil {
		return models.Companies{}, err // Devuelve el error si no se encuentra la compañía.
	}

	// Si la compañía existe, procede con la actualización.
	if err := database.DB.Model(&existing).Updates(company).Error; err != nil {
		return models.Companies{}, err
	}
	return existing, nil
}

// DeleteCompany elimina una compañía por su ID.
func DeleteCompany(id uint) error {
	var company models.Companies
	// Primero, verifica si la compañía existe.
	if err := database.DB.First(&company, id).Error; err != nil {
		return err // Devuelve el error si la compañía no se encuentra.
	}

	// Si la compañía existe, procede con la eliminación.
	if err := database.DB.Delete(&company).Error; err != nil {
		return err
	}

	return nil
}
