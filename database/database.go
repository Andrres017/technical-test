package database

import (
	"log"

	"github.com/andrres017/technical-test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Error al conectar a la base de datos:", err)
	}

	DB = db

	if err := db.AutoMigrate(&models.User{}, &models.Challenge{}, &models.Companies{}, &models.Program{}, &models.ProgramParticipant{}); err != nil {
		log.Panic("Error al realizar la migración:", err)
	}
}
