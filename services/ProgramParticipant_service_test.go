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

func TestFetchProgramParticipants(t *testing.T) {
	db := setupDatabaseForProgramParticipants(t)
	defer db.Migrator().DropTable(&models.ProgramParticipant{})

	_, _ = CreateProgramParticipant(models.ProgramParticipant{ProgramID: 1, ParticipantID: 1, ParticipantType: models.UserType})
	_, _ = CreateProgramParticipant(models.ProgramParticipant{ProgramID: 1, ParticipantID: 2, ParticipantType: models.ChallengeType})

	pp, total, err := FetchProgramParticipants(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, pp, 2)
}

func TestGetProgramParticipantByID(t *testing.T) {
	db := setupDatabaseForProgramParticipants(t)
	defer db.Migrator().DropTable(&models.ProgramParticipant{})

	newPP, _ := CreateProgramParticipant(models.ProgramParticipant{ProgramID: 1, ParticipantID: 1, ParticipantType: models.UserType})

	foundPP, err := GetProgramParticipantByID(newPP.ID)

	assert.NoError(t, err)
	assert.Equal(t, newPP.ParticipantID, foundPP.ParticipantID)
}

func TestUpdateProgramParticipant(t *testing.T) {
	db := setupDatabaseForProgramParticipants(t)
	defer db.Migrator().DropTable(&models.ProgramParticipant{})

	pp, _ := CreateProgramParticipant(models.ProgramParticipant{ProgramID: 1, ParticipantID: 1, ParticipantType: models.UserType})

	pp.ParticipantID = 2
	updatedPP, err := UpdateProgramParticipant(pp, pp.ID)

	assert.NoError(t, err)
	assert.Equal(t, uint(2), updatedPP.ParticipantID)
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

func TestCreateProgramParticipantWithError(t *testing.T) {
	db := setupDatabaseForProgramParticipants(t)
	defer db.Migrator().DropTable(&models.ProgramParticipant{})

	// Intenta crear una asociación de participante de programa sin especificar ProgramID o ParticipantID
	pp := models.ProgramParticipant{
		ParticipantType: models.UserType,
	}

	_, err := CreateProgramParticipant(pp)

	// Se espera un error porque los campos ProgramID y ParticipantID son requeridos y están vacíos.
	assert.Error(t, err)
}

func TestGetProgramParticipantByIDWithError(t *testing.T) {
	db := setupDatabaseForProgramParticipants(t)
	defer db.Migrator().DropTable(&models.ProgramParticipant{})

	// Intenta buscar una asociación de participante de programa con un ID que no existe.
	_, err := GetProgramParticipantByID(999999)

	// Se espera un error porque el ID no existe en la base de datos.
	assert.Error(t, err)
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
