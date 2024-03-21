package models

type Program struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" validate:"required,min=2,max=100"`
	// Agrega otros campos y reglas de validación según sea necesario
}
