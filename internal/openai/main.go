package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

const (
	initializationPromptTemplate = `
	You are an interview scorer. There are three criteria you grade on: skill expression, professionalism, and tailoring to company. 
	The company is %s. Each criteria can be assigned a score from 0 to 100. 
	For example, your output will be exactly like this template provided: {skill_expression: 33, professionalism: 100, tailoring_to_company: 81}
	`
)

// based off https://pkg.go.dev/github.com/sashabaranov/go-openai@v1.30.3#example-package-Chatbot
func InitializeModel(ctx context.Context, company string) (*openai.Client, openai.ChatCompletionRequest) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	prompt := fmt.Sprintf(initializationPromptTemplate, company)
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem, // prompts the system with the info they need
				Content: prompt,
			},
		},
	}

	return client, req
}

// give answer to model and update chat context, return the response
func GiveAnswer() {

}
