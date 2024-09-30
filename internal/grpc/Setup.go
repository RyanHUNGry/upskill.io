package api

import (
	"context"
	"fmt"
	"io"

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

type AnswerResponse = map[string]map[string]string

func (s *interviewServiceServer) CreateAnswer(stream InterviewService_CreateAnswerServer) error {
	createAnswerRequests := make([]AnswerResponse, 0, 10) // store responses in memory for processing when stream closes
	firstCreateAnswerRequest, err := stream.Recv()

	interviewScorer := scorer.InitializeModel(s.ctx, firstCreateAnswerRequest.CompanyName) // take the first message to get the company name
	firstResponse := interviewScorer.GiveAnswer(firstCreateAnswerRequest.Question, firstCreateAnswerRequest.Answer)

	createAnswerRequests = append(createAnswerRequests, firstResponse)

	if err == io.EOF {
		return stream.SendAndClose(nil)
	}
	if err != nil {
		return err
	}

	for {
		createAnswerRequest, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(nil)
		}
		if err != nil {
			return err
		}

		// var answer string = createAnswerRequest.Answer
		// var sessionId []byte = createAnswerRequest.SessionId
		// var interviewId []byte = createAnswerRequest.InterviewId
		// var userId []byte = createAnswerRequest.UserId
		// var questionIdx int32 = createAnswerRequest.QuestionIdx
		// var question string = createAnswerRequest.Question

		response := interviewScorer.GiveAnswer(createAnswerRequest.Question, createAnswerRequest.Answer)
		fmt.Println(response)
		createAnswerRequests = append(createAnswerRequests, response)
	}
}

func CreateInterviewServiceServer(ctx context.Context, db *cassandra.InterviewServiceDatabase) InterviewServiceServer {
	return &interviewServiceServer{ctx: ctx, db: db}
}
