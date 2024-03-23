package models

// ParticipantType define los tipos permitidos de participantes.
type ParticipantType string

const (
	// Definimos las constantes para los tipos de participantes.
	UserType      ParticipantType = "user"
	ChallengeType ParticipantType = "challenge"
	CompanyType   ParticipantType = "company"
)

// ProgramParticipant representa la asociación entre un programa y sus participantes.
type ProgramParticipant struct {
	ID              uint            `json:"id" gorm:"primaryKey"`
	ProgramID       uint            `json:"program_id" gorm:"not null"`                         // Referencia al ID del programa
	ParticipantID   uint            `json:"participant_id" gorm:"not null"`                     // ID del participante
	ParticipantType ParticipantType `json:"participant_type" gorm:"type:varchar(100);not null"` // Tipo de participante
}

type ProgramParticipantDetail struct {
	ProgramParticipant ProgramParticipant // La asociación de participante de programa original
	Program            Program            `json:"program"`            // Información del programa asociado
	ParticipantDetail  interface{}        `json:"participant_detail"` // Detalles específicos del participante
}
