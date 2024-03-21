// utils/validators.go

package utils

import "github.com/andrres017/technical-test/models"

// IsParticipantTypeValid valida si el tipo de participante es uno permitido.
func IsParticipantTypeValid(pt models.ParticipantType) bool {
	switch pt {
	case models.UserType, models.ChallengeType, models.CompanyType:
		return true
	default:
		return false
	}
}
