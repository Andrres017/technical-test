package services

import (
	"testing"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabaseForProgramParticipants(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.AutoMigrate(&models.ProgramParticipant{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	database.DB = db
	return db
}

func TestCreateProgramParticipant(t *testing.T) {
	db := setupDatabaseForProgramParticipants(t)
	defer db.Migrator().DropTable(&models.ProgramParticipant{})

	pp := models.ProgramParticipant{
		ProgramID:       1,
		ParticipantID:   1,
		ParticipantType: models.UserType,
	}

	result, err := CreateProgramParticipant(pp)

	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
}

func TestDeleteProgramParticipant(t *testing.T) {
	db := setupDatabaseForProgramParticipants(t)
	defer db.Migrator().DropTable(&models.ProgramParticipant{})

	pp, _ := CreateProgramParticipant(models.ProgramParticipant{ProgramID: 1, ParticipantID: 1, ParticipantType: models.UserType})

	err := DeleteProgramParticipant(pp.ID)

	assert.NoError(t, err)

	_, err = GetProgramParticipantByID(pp.ID)
	assert.Error(t, err) // Expect an error because the program participant should no longer exist
}

func TestUpdateProgramParticipantWithError(t *testing.T) {
	db := setupDatabaseForProgramParticipants(t)
	defer db.Migrator().DropTable(&models.ProgramParticipant{})

	// Intenta actualizar una asociación de participante de programa inexistente.
	pp := models.ProgramParticipant{ID: 999999, ParticipantType: models.UserType}

	_, err := UpdateProgramParticipant(pp, pp.ID)

	// Se espera un error porque la asociación con el ID especificado no existe.
	assert.Error(t, err)
}

func TestDeleteProgramParticipantWithError(t *testing.T) {
	db := setupDatabaseForProgramParticipants(t)
	defer db.Migrator().DropTable(&models.ProgramParticipant{})

	// Intenta eliminar una asociación de participante de programa que no existe.
	err := DeleteProgramParticipant(999999)

	// Se espera un error porque la asociación con el ID especificado no existe.
	assert.Error(t, err)
}
