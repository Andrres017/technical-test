package services

import (
	"testing"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Challenge{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	database.DB = db // Use this db for operations
	return db
}

func TestCreateChallenge(t *testing.T) {
	db := setupDatabase(t)
	defer db.Migrator().DropTable(&models.Challenge{})

	challenge := models.Challenge{Name: "Test Challenge"}

	result, err := CreateChallenge(challenge)

	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
	assert.Equal(t, "Test Challenge", result.Name)
}

func TestFetchChallenges(t *testing.T) {
	db := setupDatabase(t)
	defer db.Migrator().DropTable(&models.Challenge{})

	_, err := CreateChallenge(models.Challenge{Name: "Test Challenge 1"})
	assert.NoError(t, err)
	_, err = CreateChallenge(models.Challenge{Name: "Test Challenge 2"})
	assert.NoError(t, err)

	challenges, total, err := FetchChallenges(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, challenges, 2)
}

func TestGetChallengeByID(t *testing.T) {
	db := setupDatabase(t)
	defer db.Migrator().DropTable(&models.Challenge{})

	newChallenge, err := CreateChallenge(models.Challenge{Name: "Test Challenge"})
	assert.NoError(t, err)

	challenge, err := GetChallengeByID(newChallenge.ID)

	assert.NoError(t, err)
	assert.Equal(t, newChallenge.Name, challenge.Name)
}

func TestUpdateChallenge(t *testing.T) {
	db := setupDatabase(t)
	defer db.Migrator().DropTable(&models.Challenge{})

	newChallenge, err := CreateChallenge(models.Challenge{Name: "Old Name"})
	assert.NoError(t, err)

	newChallenge.Name = "New Name"
	updatedChallenge, err := UpdateChallenge(newChallenge, newChallenge.ID)

	assert.NoError(t, err)
	assert.Equal(t, "New Name", updatedChallenge.Name)
}

func TestDeleteChallenge(t *testing.T) {
	db := setupDatabase(t)
	defer db.Migrator().DropTable(&models.Challenge{})

	challenge, err := CreateChallenge(models.Challenge{Name: "Test Challenge"})
	assert.NoError(t, err)

	err = DeleteChallenge(challenge.ID)

	assert.NoError(t, err)

	_, err = GetChallengeByID(challenge.ID)
	assert.Error(t, err) // Expect an error because the challenge should no longer exist
}

func TestCreateChallengeWithError(t *testing.T) {
	db := setupDatabase(t)
	defer db.Migrator().DropTable(&models.Challenge{})

	// Intenta crear un Challenge sin nombre, asumiendo que el campo 'Name' es requerido.
	challenge := models.Challenge{}

	_, err := CreateChallenge(challenge)

	// Se espera un error porque el campo 'Name' es requerido y está vacío.
	assert.Error(t, err)
}

func TestGetChallengeByIDWithError(t *testing.T) {
	db := setupDatabase(t)
	defer db.Migrator().DropTable(&models.Challenge{})

	// Intenta buscar un Challenge con un ID que no existe.
	_, err := GetChallengeByID(999999)

	// Se espera un error porque el ID no existe en la base de datos.
	assert.Error(t, err)
}

func TestUpdateChallengeWithError(t *testing.T) {
	db := setupDatabase(t)
	defer db.Migrator().DropTable(&models.Challenge{})

	// Intenta actualizar un Challenge inexistente.
	challenge := models.Challenge{ID: 999999, Name: "Nonexistent Challenge"}

	_, err := UpdateChallenge(challenge, challenge.ID)

	// Se espera un error porque el Challenge con el ID especificado no existe.
	assert.Error(t, err)
}

func TestDeleteChallengeWithError(t *testing.T) {
	db := setupDatabase(t)
	defer db.Migrator().DropTable(&models.Challenge{})

	// Intenta eliminar un Challenge que no existe.
	err := DeleteChallenge(999999)

	// Se espera un error porque el Challenge con el ID especificado no existe.
	assert.Error(t, err)
}
