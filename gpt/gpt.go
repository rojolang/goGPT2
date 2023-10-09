// gp4.go
// author: rjjm94 on 10/09/2023
// functions added: GetGPT4Response
// variables: GPT4
// description: This file contains the functionality to generate a response from the GPT-4 API.

package gpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

// GetGPT4Response generates a response from the GPT-4 API based on the given prompt, system message, max tokens, and temperature
func GetGPT4Response(prompt string, systemMessage string, maxTokens int, temperature float32) (string, error) {
	// Create a new OpenAI client
	client := openai.NewClient(os.Getenv("OPENAI_SECRET_KEY"))

	// Create a chat completion request
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemMessage,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens:   maxTokens,
			Temperature: temperature,
		},
	)

	if err != nil {
		return "", err
	}

	// Return the content of the first message in the response
	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}
	return "", fmt.Errorf("no choices in the response")
}
