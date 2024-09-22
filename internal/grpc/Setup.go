package api

import (
	"context"
	"fmt"
	"io"
)

type interviewServiceServer struct {
	UnimplementedInterviewServiceServer // forward compatability, and implementing the empty method
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

		// For each
		fmt.Println(createAnswerRequest)
		// process each createAnswerRequest
		fmt.Println(createAnswerRequest.Answer)
		fmt.Println(createAnswerRequest.SessionId)
	}
}

func CreateInterviewServiceServer() InterviewServiceServer {
	return &interviewServiceServer{}
}
