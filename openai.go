package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Replace "YOUR_API_KEY" with your actual OpenAI API key.
const apiKey = "YOUR_API_KEY"
const openAIURL = "https://api.openai.com/v1/engines/davinci-codex/completions"

// RequestData represents the data sent in the API request.
type RequestData struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

// ResponseData represents the data received in the API response.
type ResponseData struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// OpenAIAPI is a struct to make OpenAI API requests.
type OpenAIAPI struct {
	client *http.Client
}

// NewOpenAIAPI creates a new instance of OpenAIAPI with a custom HTTP client.
func NewOpenAIAPI() *OpenAIAPI {
	return &OpenAIAPI{
		client: &http.Client{},
	}
}

// Complete sends a text completion request to the OpenAI API and returns the response.
func (oa *OpenAIAPI) Complete(prompt string, maxTokens int) (string, error) {
	data := RequestData{
		Prompt:    prompt,
		MaxTokens: maxTokens,
	}

	requestJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	requestBody := bytes.NewReader(requestJSON)

	req, err := http.NewRequest("POST", openAIURL, requestBody)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := oa.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseData ResponseData
	err = json.Unmarshal(responseJSON, &responseData)
	if err != nil {
		return "", err
	}

	if len(responseData.Choices) > 0 {
		return responseData.Choices[0].Text, nil
	}

	return "", fmt.Errorf("no response data")
}

// func main() {
// 	api := NewOpenAIAPI()

// 	prompt := "Once upon a time, there was a "
// 	maxTokens := 50

// 	result, err := api.Complete(prompt, maxTokens)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Generated Text:", result)
// }
