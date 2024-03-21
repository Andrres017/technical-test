package services

import (
	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
)

// CreateCompany crea una nueva compañía en la base de datos.
func CreateCompany(company models.Companies) (models.Companies, error) {
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

// UpdateCompany actualiza una compañía existente.
func UpdateCompany(company models.Companies, id uint) (models.Companies, error) {
	if err := database.DB.Model(&models.Companies{}).Where("id = ?", id).Updates(company).Error; err != nil {
		return models.Companies{}, err
	}
	return company, nil
}

// DeleteCompany elimina una compañía por su ID.
func DeleteCompany(id uint) error {
	result := database.DB.Delete(&models.Companies{}, id)
	return result.Error
}
