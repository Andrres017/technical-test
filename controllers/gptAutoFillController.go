package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/andrres017/technical-test/models"
	"github.com/andrres017/technical-test/services"
	"github.com/andrres017/technical-test/utils"
	"github.com/labstack/echo/v4"
)

// GenerateFakeData genera datos falsos para el modelo especificado
func GenerateFakeData(model string) (string, error) {
	prompt := fmt.Sprintf("Generate a pure JSON array with 10 fake data entries for the %s model. Each entry should include a 'name' field. Do not include any extra characters or words, just the JSON array.", model)
	return services.GPTAutoFill(prompt)
}

func cleanJSONInput(input string) string {
	// Eliminar los backticks y ajustes de formato que rodean al verdadero JSON
	cleanedInput := strings.Replace(input, "```json", "", -1)
	cleanedInput = strings.Replace(cleanedInput, "```", "", -1)
	cleanedInput = strings.TrimSpace(cleanedInput) // Elimina espacios en blanco adicionales
	return cleanedInput
}

// HandleGPTAutoFill processes requests to auto-fill database tables with fake data
func HandleGPTAutoFill(c echo.Context) error {
	tables := []string{"Challenge", "Companies", "Program", "User"}

	var wg sync.WaitGroup

	for _, model := range tables {
		wg.Add(1)
		go func(model string) {
			defer wg.Done()

			result, err := GenerateFakeData(model)
			fmt.Printf("data Generate %s", result)
			if err != nil {
				fmt.Printf("Error generating fake data for model %s: %v\n", model, err)
				return
			}

			cleanedResult := cleanJSONInput(result)

			// Aquí manejaremos cada caso según el modelo, parseando y creando adecuadamente.
			switch model {
			case "Challenge":
				var challenges []models.Challenge
				err = json.Unmarshal([]byte(cleanedResult), &challenges)
				if err != nil {
					fmt.Printf("Error parsing JSON for model %s: %v\n", model, err)
					return
				}

				for _, challenge := range challenges {
					_, err := services.CreateChallenge(challenge)
					if err != nil {
						fmt.Printf("Error creating Challenge: %v\n", err)
					}
				}
			case "Companies":
				var companies []models.Companies
				err = json.Unmarshal([]byte(cleanedResult), &companies)
				if err != nil {
					fmt.Printf("Error parsing JSON for model %s: %v\n", model, err)
					return
				}

				for _, companies := range companies {
					_, err := services.CreateCompany(companies)
					if err != nil {
						fmt.Printf("Error creating Challenge: %v\n", err)
					}
				}
			case "Program":
				var program []models.Program
				err = json.Unmarshal([]byte(cleanedResult), &program)
				if err != nil {
					fmt.Printf("Error parsing JSON for model %s: %v\n", model, err)
					return
				}

				for _, program := range program {
					_, err := services.CreateProgram(program)
					if err != nil {
						fmt.Printf("Error creating Challenge: %v\n", err)
					}
				}
			case "User":
				var user []models.User
				err = json.Unmarshal([]byte(cleanedResult), &user)
				if err != nil {
					fmt.Printf("Error parsing JSON for model %s: %v\n", model, err)
					return
				}

				for _, user := range user {
					_, err := services.CreateUser(user)
					if err != nil {
						fmt.Printf("Error creating Challenge: %v\n", err)
					}
				}
			default:
				fmt.Printf("Model %s not recognized\n", model)
			}
		}(model)
	}

	wg.Wait()
	fmt.Println("All data generation and saving complete.")

	return utils.ApiResponse(c, http.StatusOK, "Success", "Fake data generated", nil)
}
