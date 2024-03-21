package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// GPTAutoFill sends a request to the GPT-3.5 API using the chat completions endpoint and returns the GPT response message.
func GPTAutoFill(prompt string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions" // Updated endpoint for chat models
	token := os.Getenv("OPENAI_API_KEY")

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo", // Ensure you use the correct chat model
		"messages": []map[string]string{ // Format for chat completions
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": prompt},
		},
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 && len(response.Choices[0].Message.Content) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", nil // O devuelve un error si prefieres manejar la situación de no tener respuesta.
}
