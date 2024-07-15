package server

import (
	"context"
	pb "interview-service/api"

	"go.mongodb.org/mongo-driver/mongo"
)

type PromptServiceServer struct {
	db *mongo.Client
	pb.UnimplementedPromptServiceServer
}

type AnswerServiceServer struct {
	db *mongo.Client
	pb.UnimplementedAnswerServiceServer
}

// INTEFACE IMPLEMENTATION START
func (s *PromptServiceServer) GetPrompt(ctx context.Context, req *pb.PromptRequest) (*pb.PromptResponse, error) {
	return nil, nil
}

func (s *AnswerServiceServer) PostAnswer(ctx context.Context, req *pb.AnswerRequest) (*pb.AnswerResponse, error) {
	return nil, nil
}

// INTERFACE IMPLEMENTATION END

func (s *PromptServiceServer) SetDatabase(db *mongo.Client) {
	s.db = db
}

func (s *AnswerServiceServer) SetDatabase(db *mongo.Client) {
	s.db = db
}
