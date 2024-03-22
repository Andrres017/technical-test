// database/interface.go (nuevo archivo, si decides ir por este camino)

package database

import (
	"github.com/andrres017/technical-test/models"
	"gorm.io/gorm"
)

type ChallengeDB interface {
	Create(challenge *models.Challenge) *gorm.DB
}
