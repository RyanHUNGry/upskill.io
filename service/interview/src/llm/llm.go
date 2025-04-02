package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/responses"
	"github.com/openai/openai-go/shared"
)

// THE client is the new Response API rather than the original Chat Completion API
type Model struct {
	Client openai.Client
	Ctx    context.Context
}

func InitializeModel(ctx context.Context) *Model {
	apiKey := os.Getenv("OPENAI_API_DEVKEY")
	clientOpts := []option.RequestOption{option.WithAPIKey(apiKey)}
	client := openai.NewClient(clientOpts...)
	return &Model{Client: client, Ctx: ctx}
}

func (m Model) GenerateFeedback(question string, answer string) (string, error) {
	body := responses.ResponseNewParams{
		Model: shared.ChatModelGPT4oMini,
		Input: responses.ResponseNewParamsInputUnion{
			OfString: param.Opt[string]{Value: question},
		},
	}

	res, err := m.Client.Responses.New(m.Ctx, body)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	return "YES", nil
}

func (m Model) GenerateQuestion() (string, error) {
	body := responses.ResponseNewParams{
		Model: shared.ChatModelGPT4oMini,
		Input: responses.ResponseNewParamsInputUnion{
			OfString: param.Opt[string]{Value: question},
		},
	}

	res, err := m.Client.Responses.New(m.Ctx, body)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	return "YES", nil
}
