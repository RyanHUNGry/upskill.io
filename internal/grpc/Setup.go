package api

import (
	"context"
	"fmt"
	"io"

	"github.com/sashabaranov/go-openai"
	"ryhung.upskill.io/internal/cassandra"
	scorer "ryhung.upskill.io/internal/openai"
)

type interviewServiceServer struct {
	UnimplementedInterviewServiceServer // forward compatability, and implementing the empty method
	ctx                                 context.Context
	db                                  *cassandra.InterviewServiceDatabase
}

func (s *interviewServiceServer) CreateInterview(context.Context, *CreateInterviewRequest) (*GetInterview, error) {
	return nil, nil
}

func (s *interviewServiceServer) CreateAnswer(stream InterviewService_CreateAnswerServer) error {
	for {
		createAnswerRequest, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(new(GetAnswerScores))
		}
		if err != nil {
			return err
		}

		// Build initial prompt
		client, req := scorer.InitializeModel(s.ctx, "Jane Street")

		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: createAnswerRequest.Answer,
		})

		resp, err := client.CreateChatCompletion(s.ctx, req)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}
}

func CreateInterviewServiceServer(ctx context.Context, db *cassandra.InterviewServiceDatabase) InterviewServiceServer {
	return &interviewServiceServer{ctx: ctx, db: db}
}
