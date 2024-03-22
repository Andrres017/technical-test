package services

import (
	"testing"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabaseForCompanies(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate the schema for Companies
	if err := db.AutoMigrate(&models.Companies{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	database.DB = db // Use this db for operations
	return db
}

func TestCreateCompany(t *testing.T) {
	db := setupDatabaseForCompanies(t)
	defer db.Migrator().DropTable(&models.Companies{})

	company := models.Companies{Name: "Test Company"}

	result, err := CreateCompany(company)

	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
	assert.Equal(t, "Test Company", result.Name)
}

func TestFetchCompanies(t *testing.T) {
	db := setupDatabaseForCompanies(t)
	defer db.Migrator().DropTable(&models.Companies{})

	_, err := CreateCompany(models.Companies{Name: "Test Company 1"})
	assert.NoError(t, err)
	_, err = CreateCompany(models.Companies{Name: "Test Company 2"})
	assert.NoError(t, err)

	companies, total, err := FetchCompanies(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, companies, 2)
}

func TestGetCompanyByID(t *testing.T) {
	db := setupDatabaseForCompanies(t)
	defer db.Migrator().DropTable(&models.Companies{})

	newCompany, err := CreateCompany(models.Companies{Name: "Test Company"})
	assert.NoError(t, err)

	company, err := GetCompanyByID(newCompany.ID)

	assert.NoError(t, err)
	assert.Equal(t, "Test Company", company.Name)
}

func TestUpdateCompany(t *testing.T) {
	db := setupDatabaseForCompanies(t)
	defer db.Migrator().DropTable(&models.Companies{})

	newCompany, err := CreateCompany(models.Companies{Name: "Old Name"})
	assert.NoError(t, err)

	newCompany.Name = "New Name"
	updatedCompany, err := UpdateCompany(newCompany, newCompany.ID)

	assert.NoError(t, err)
	assert.Equal(t, "New Name", updatedCompany.Name)
}

func TestDeleteCompany(t *testing.T) {
	db := setupDatabaseForCompanies(t)
	defer db.Migrator().DropTable(&models.Companies{})

	company, err := CreateCompany(models.Companies{Name: "Test Company"})
	assert.NoError(t, err)

	err = DeleteCompany(company.ID)

	assert.NoError(t, err)

	_, err = GetCompanyByID(company.ID)
	assert.Error(t, err) // Expect an error because the company should no longer exist
}

func TestCreateCompanyWithError(t *testing.T) {
	db := setupDatabaseForCompanies(t)
	defer db.Migrator().DropTable(&models.Companies{})

	// Intenta crear una compañía sin nombre (asumiendo que el campo 'Name' es requerido y no puede estar vacío).
	company := models.Companies{}

	_, err := CreateCompany(company)

	// Se espera un error porque el campo 'Name' es requerido.
	assert.Error(t, err)
}

func TestGetCompanyByIDWithError(t *testing.T) {
	db := setupDatabaseForCompanies(t)
	defer db.Migrator().DropTable(&models.Companies{})

	// Intenta buscar una compañía con un ID que no existe.
	_, err := GetCompanyByID(999999)

	// Se espera un error porque el ID no existe en la base de datos.
	assert.Error(t, err)
}

func TestUpdateCompanyWithError(t *testing.T) {
	db := setupDatabaseForCompanies(t)
	defer db.Migrator().DropTable(&models.Companies{})

	// Prepara una compañía con un ID que no existe.
	company := models.Companies{ID: 999999, Name: "Nonexistent Company"}

	// Intenta actualizar la compañía.
	_, err := UpdateCompany(company, company.ID)

	// Se espera un error porque la compañía con el ID especificado no existe.
	assert.Error(t, err)
}

func TestDeleteCompanyWithError(t *testing.T) {
	db := setupDatabaseForCompanies(t)
	defer db.Migrator().DropTable(&models.Companies{})

	// Intenta eliminar una compañía con un ID que no existe.
	err := DeleteCompany(999999)

	// Se espera un error porque la compañía con el ID especificado no existe.
	assert.Error(t, err)
}
