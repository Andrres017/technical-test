package models

type Companies struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" validate:"required,min=2,max=100"`
	// Agrega otros campos y reglas de validación según sea necesario
}
