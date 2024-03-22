package usecases

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/andrres017/technical-test/models"
	"github.com/andrres017/technical-test/services"
)

type GPTAutoFillUseCase interface {
	GenerateAndSaveFakeData(model string) error
}

type gptAutoFillUseCaseImpl struct{}

func NewGPTAutoFillUseCase() GPTAutoFillUseCase {
	return &gptAutoFillUseCaseImpl{}
}

func (uc *gptAutoFillUseCaseImpl) GenerateAndSaveFakeData(model string) error {
	prompt := fmt.Sprintf("Generate a pure JSON array with 10 fake data entries for the %s model. Each entry should include a 'name' field. Do not include any extra characters or words, just the JSON array.", model)
	result, err := services.GPTAutoFill(prompt)
	if err != nil {
		return fmt.Errorf("error generating fake data for model %s: %w", model, err)
	}

	cleanedResult := cleanJSONInput(result)

	switch model {
	case "Challenge":
		return parseAndCreateChallenges(cleanedResult)
	case "Companies":
		return parseAndCreateCompanies(cleanedResult)
	case "Program":
		return parseAndCreatePrograms(cleanedResult)
	case "User":
		return parseAndCreateUsers(cleanedResult)
	default:
		return fmt.Errorf("model %s not recognized", model)
	}
}

func cleanJSONInput(input string) string {
	cleanedInput := strings.Replace(input, "```json", "", -1)
	cleanedInput = strings.Replace(cleanedInput, "```", "", -1)
	cleanedInput = strings.TrimSpace(cleanedInput)
	return cleanedInput
}

func parseAndCreateChallenges(jsonData string) error {
	var challenges []models.Challenge
	if err := json.Unmarshal([]byte(jsonData), &challenges); err != nil {
		return fmt.Errorf("error parsing JSON for Challenges: %w", err)
	}

	for _, challenge := range challenges {
		if _, err := services.CreateChallenge(challenge); err != nil {
			return fmt.Errorf("error creating Challenge: %w", err)
		}
	}

	return nil
}

func parseAndCreateCompanies(jsonData string) error {
	// Aquí debes añadir la lógica para deserializar el JSON a []models.Companies y luego crearlos en la base de datos.
	var companies []models.Companies
	if err := json.Unmarshal([]byte(jsonData), &companies); err != nil {
		return fmt.Errorf("error parsing JSON for Companies: %w", err)
	}

	for _, companies := range companies {
		if _, err := services.CreateCompany(companies); err != nil {
			return fmt.Errorf("error creating Companies: %w", err)
		}
	}

	return nil
}

func parseAndCreatePrograms(jsonData string) error {
	// Similar a parseAndCreateChallenges, implementa la lógica para Programs.
	var program []models.Program
	if err := json.Unmarshal([]byte(jsonData), &program); err != nil {
		return fmt.Errorf("error parsing JSON for program: %w", err)
	}

	for _, program := range program {
		if _, err := services.CreateProgram(program); err != nil {
			return fmt.Errorf("error creating program: %w", err)
		}
	}
	return nil
}

func parseAndCreateUsers(jsonData string) error {
	// Implementa la lógica de deserialización y creación para Users.
	var user []models.User
	if err := json.Unmarshal([]byte(jsonData), &user); err != nil {
		return fmt.Errorf("error parsing JSON for user: %w", err)
	}

	for _, user := range user {
		if _, err := services.CreateUser(user); err != nil {
			return fmt.Errorf("error creating user: %w", err)
		}
	}

	return nil
}
