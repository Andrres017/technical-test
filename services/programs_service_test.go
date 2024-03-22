package services

import (
	"testing"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabaseForPrograms(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.AutoMigrate(&models.Program{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	database.DB = db
	return db
}

func TestCreateProgram(t *testing.T) {
	db := setupDatabaseForPrograms(t)
	defer db.Migrator().DropTable(&models.Program{})

	program := models.Program{Name: "New Program"}

	result, err := CreateProgram(program)

	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
	assert.Equal(t, "New Program", result.Name)
}

func TestFetchPrograms(t *testing.T) {
	db := setupDatabaseForPrograms(t)
	defer db.Migrator().DropTable(&models.Program{})

	_, _ = CreateProgram(models.Program{Name: "Program 1"})
	_, _ = CreateProgram(models.Program{Name: "Program 2"})

	programs, total, err := FetchPrograms(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, programs, 2)
}

func TestGetProgramByID(t *testing.T) {
	db := setupDatabaseForPrograms(t)
	defer db.Migrator().DropTable(&models.Program{})

	newProgram, _ := CreateProgram(models.Program{Name: "Program By ID"})

	foundProgram, err := GetProgramByID(newProgram.ID)

	assert.NoError(t, err)
	assert.Equal(t, "Program By ID", foundProgram.Name)
}

func TestUpdateProgram(t *testing.T) {
	db := setupDatabaseForPrograms(t)
	defer db.Migrator().DropTable(&models.Program{})

	program, _ := CreateProgram(models.Program{Name: "Initial Name"})

	program.Name = "Updated Name"
	updatedProgram, err := UpdateProgram(program, program.ID)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedProgram.Name)
}

func TestDeleteProgram(t *testing.T) {
	db := setupDatabaseForPrograms(t)
	defer db.Migrator().DropTable(&models.Program{})

	program, _ := CreateProgram(models.Program{Name: "To Be Deleted"})

	err := DeleteProgram(program.ID)

	assert.NoError(t, err)

	_, err = GetProgramByID(program.ID)
	assert.Error(t, err) // Expect an error because the program should no longer exist
}

func TestCreateProgramWithError(t *testing.T) {
	db := setupDatabaseForPrograms(t)
	defer db.Migrator().DropTable(&models.Program{})

	// Intenta crear un Program sin nombre, asumiendo que el campo 'Name' es requerido.
	program := models.Program{}

	_, err := CreateProgram(program)

	// Se espera un error porque el campo 'Name' es requerido y está vacío.
	assert.Error(t, err)
}

func TestGetProgramByIDWithError(t *testing.T) {
	db := setupDatabaseForPrograms(t)
	defer db.Migrator().DropTable(&models.Program{})

	// Intenta buscar un Program con un ID que no existe.
	_, err := GetProgramByID(999999)

	// Se espera un error porque el ID no existe en la base de datos.
	assert.Error(t, err)
}

func TestUpdateProgramWithError(t *testing.T) {
	db := setupDatabaseForPrograms(t)
	defer db.Migrator().DropTable(&models.Program{})

	// Intenta actualizar un Program inexistente.
	program := models.Program{ID: 999999, Name: "Nonexistent Program"}

	_, err := UpdateProgram(program, program.ID)

	// Se espera un error porque el Program con el ID especificado no existe.
	assert.Error(t, err)
}

func TestDeleteProgramWithError(t *testing.T) {
	db := setupDatabaseForPrograms(t)
	defer db.Migrator().DropTable(&models.Program{})

	// Intenta eliminar un Program que no existe.
	err := DeleteProgram(999999)

	// Se espera un error porque el Program con el ID especificado no existe.
	assert.Error(t, err)
}
