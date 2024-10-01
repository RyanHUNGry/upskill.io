package api

import (
	"context"
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

type AnswerResponsePair struct {
	answerResponse AnswerResponse // this is a map, which is referenced by pointer under the hood, so this struct doesn't need to be pointer referenced
	answerRequest  *CreateAnswerRequest
}
type AnswerResponse = map[string]map[string]interface{}

func (s *interviewServiceServer) CreateAnswer(stream InterviewService_CreateAnswerServer) error {
	createAnswerRequests := make([]AnswerResponsePair, 0, 10)                              // store responses in memory for processing when stream closes. Handler activates with a gauranteed first message.
	firstCreateAnswerRequest, err := stream.Recv()                                         // blocks until new line is received, then sequentially in current execution thread
	interviewScorer := scorer.InitializeModel(s.ctx, firstCreateAnswerRequest.CompanyName) // take the first message to get the company name
	firstResponse := interviewScorer.GiveAnswer(firstCreateAnswerRequest.Question, firstCreateAnswerRequest.Answer)

	createAnswerRequests = append(createAnswerRequests, AnswerResponsePair{firstResponse, firstCreateAnswerRequest})

	if err == io.EOF {
		getAnswerScores, err := answerAggregator(createAnswerRequests)
		if err != nil {
			return err
		}
		return stream.SendAndClose(getAnswerScores)
	}
	if err != nil {
		return err
	}

	for {
		createAnswerRequest, err := stream.Recv()
		if err == io.EOF {
			getAnswerScores, err := answerAggregator(createAnswerRequests)
			if err != nil {
				return err
			}
			return stream.SendAndClose(getAnswerScores)
		}
		if err != nil {
			return err
		}

		response := interviewScorer.GiveAnswer(createAnswerRequest.Question, createAnswerRequest.Answer)
		createAnswerRequests = append(createAnswerRequests, AnswerResponsePair{response, createAnswerRequest})
	}
}

func CreateInterviewServiceServer(ctx context.Context, db *cassandra.InterviewServiceDatabase) InterviewServiceServer {
	return &interviewServiceServer{ctx: ctx, db: db}
}
