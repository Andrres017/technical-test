package utils

import (
	"github.com/labstack/echo/v4"
)

// ResponseFormat define la estructura de las respuestas de la API
type ResponseFormat struct {
	Title   string      `json:"title,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ApiResponse envía respuestas estandarizadas
func ApiResponse(c echo.Context, statusCode int, title string, message string, data interface{}) error {
	response := ResponseFormat{
		Title:   title,
		Message: message,
		Data:    data,
	}
	return c.JSON(statusCode, response)
}
