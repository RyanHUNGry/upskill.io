package scorer

/*
Structured input:
type String

"Question: What is your greatest strength? Answer: My greatest strength is my ability to communicate."

Structured output:
type JSON object
{
	"skill_expression": {
		"score": 33,
		"reason": "The candidate was able to communicate their greatest strength effectively."
	},
	"professionalism": {
		"score": 100,
		"reason": "The candidate was professional in their response."
	},
}
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type InterviewScorer struct {
	client   *openai.Client
	ctx      context.Context
	reqChain *openai.ChatCompletionRequest
}

var client *openai.Client

const (
	initializationPromptTemplate = `
	You are an interview scorer. There are three criteria you grade on: skill expression, professionalism, and tailoring to company. 
	The interviewing company is %s. You will receive each question followed by its reponse. Each criteria must be assigned a score from 0 to 100. There must be a reason given for each score.
	`
)

const (
	answerQuestionPromptTemplate = `
	Question: %s
	Answer: %s
	`
)

// based off https://platform.openai.com/docs/guides/function-calling
// sashabaranov/go-openai extends the API above in golang pretty one-to-one
var schema *openai.FunctionDefinition = &openai.FunctionDefinition{
	Name:        "scorer",
	Description: "This is the official scorer output schema",
	Parameters: jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"skill_expression": {
				Type:        jsonschema.Object,
				Description: "The object for skill expression",
				Properties: map[string]jsonschema.Definition{
					"score": {
						Type:        jsonschema.Number,
						Description: "The score for skill expression on a scale of 0-100",
					},
					"reason": {
						Type:        jsonschema.String,
						Description: "The reason for the score for skill expression",
					},
				},
			},
			"professionalism": {
				Type:        jsonschema.Object,
				Description: "The object for professionalism",
				Properties: map[string]jsonschema.Definition{
					"score": {
						Type:        jsonschema.Number,
						Description: "The score for professionalism on a scale of 0-100",
					},
					"reason": {
						Type:        jsonschema.String,
						Description: "The reason for the score for professionalism",
					},
				},
			},
			"tailoring_to_company": {
				Type:        jsonschema.Object,
				Description: "The object for tailoring to company",
				Properties: map[string]jsonschema.Definition{
					"score": {
						Type:        jsonschema.Number,
						Description: "The score for tailoring to company on a scale of 0-100",
					},
					"reason": {
						Type:        jsonschema.String,
						Description: "The reason for the score for tailoring to company",
					},
				},
			},
		},
	},
}

// Rather than initializing a new client per gRPC request, we can initialize the model once and use it for all requests
func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	apiKey := os.Getenv("OPENAI_API_KEY")
	client = openai.NewClient(apiKey)
}

// based off https://pkg.go.dev/github.com/sashabaranov/go-openai@v1.30.3#example-package-Chatbot
func InitializeModel(ctx context.Context, company string) *InterviewScorer {
	prompt := fmt.Sprintf(initializationPromptTemplate, company)
	req := &openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem, // prompts the system with the info they need
				Content: prompt,
			},
		},
		Tools:      []openai.Tool{{Type: openai.ToolTypeFunction, Function: schema}},                                   // function name is scorer
		ToolChoice: openai.ToolChoice{Type: openai.ToolTypeFunction, Function: openai.ToolFunction{Name: schema.Name}}, // allows us to choose the scorer function as the output JSON schema
	}

	_, err := client.CreateChatCompletion(ctx, *req) // don't need response to system prompt
	if err != nil {
		panic(err)
	}

	return &InterviewScorer{client: client, ctx: ctx, reqChain: req}
}

// give answer to model and update chat context, return the unmarshalled response
func (interviewScorer *InterviewScorer) GiveAnswer(question, answer string) map[string]map[string]interface{} {
	prompt := fmt.Sprintf(answerQuestionPromptTemplate, question, answer)

	interviewScorer.reqChain.Messages = append(interviewScorer.reqChain.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})

	resp, err := interviewScorer.client.CreateChatCompletion(interviewScorer.ctx, *interviewScorer.reqChain)
	if err != nil {
		panic(err)
	}

	functionCallResultMessage := openai.ChatCompletionMessage{
		Role:       openai.ChatMessageRoleTool,
		ToolCallID: resp.Choices[0].Message.ToolCalls[0].ID,
		Content:    resp.Choices[0].Message.ToolCalls[0].Function.Arguments,
	}

	interviewScorer.reqChain.Messages = append(interviewScorer.reqChain.Messages, resp.Choices[0].Message, functionCallResultMessage)

	parsedResponse, err := interviewScorer.parseAnswer(resp)
	if err != nil {
		panic(err)
	}
	return parsedResponse
}

// parses the schema into a map
func (interviewScorer *InterviewScorer) parseAnswer(resp openai.ChatCompletionResponse) (map[string]map[string]interface{}, error) {
	responseMap := make(map[string]map[string]interface{})
	responseBytes := []byte(resp.Choices[0].Message.ToolCalls[0].Function.Arguments)
	json.Unmarshal(responseBytes, &responseMap)
	return responseMap, nil
}
