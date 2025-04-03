package llm

import (
	"context"
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

func (m Model) GenerateFeedback(role string, company string, description string, skills []string, question string, answer string, prevResponse *responses.Response) (*responses.Response, error) {
	var body responses.ResponseNewParams

	if prevResponse == nil {
		body = responses.ResponseNewParams{
			Model: shared.ChatModelGPT4oMini,
			Input: responses.ResponseNewParamsInputUnion{
				OfString: param.Opt[string]{Value: question},
			},
			Instructions: param.Opt[string]{Value: GenerateFeedbackPrompt(role, company, description, skills, answer, question)},
		}
	} else {
		body = responses.ResponseNewParams{
			Model: shared.ChatModelGPT4oMini,
			Input: responses.ResponseNewParamsInputUnion{
				OfString: param.Opt[string]{Value: question},
			},
			PreviousResponseID: param.Opt[string]{Value: prevResponse.ID},
		}
	}

	res, err := m.Client.Responses.New(m.Ctx, body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m Model) GenerateQuestion(role string, company string, description string, skills []string, prevResponse *responses.Response) (*responses.Response, error) {
	var body responses.ResponseNewParams

	if prevResponse == nil {
		body = responses.ResponseNewParams{
			Model: shared.ChatModelGPT4oMini,
			Input: responses.ResponseNewParamsInputUnion{
				OfString: param.Opt[string]{Value: "Hi, I just joined the interview and am ready to answer any questions you have for me!"},
			},
			Instructions: param.Opt[string]{Value: GenerateQuestionsPrompt(role, company, description, skills)},
		}
	} else {
		body = responses.ResponseNewParams{
			Model:              shared.ChatModelGPT4oMini,
			PreviousResponseID: param.Opt[string]{Value: prevResponse.ID},
			Instructions:       param.Opt[string]{Value: GenerateQuestionsPrompt(role, company, description, skills)},
		}
	}

	res, err := m.Client.Responses.New(m.Ctx, body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
