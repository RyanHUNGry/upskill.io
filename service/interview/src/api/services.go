package api

import (
	context "context"
	"interview/src/db"
)

type InterviewServiceServerImpl struct {
	UnimplementedInterviewServiceServer
	session *db.CassandraSession
}

func (s *InterviewServiceServerImpl) CreateInterviewTemplateCall(ctx context.Context, in *CreateInterviewTemplate) (*InterviewTemplate, error) {
	s.session.CreateInterviewTemplate(in.Company, in.Role, in.Skills, in.Description, in.Questions, in.UserId)
	return &InterviewTemplate{}, nil
}

func (s *InterviewServiceServerImpl) CreateConductedInterviewCall(ctx context.Context, in *CreateConductedInterview) (*ConductedInterview, error) {
	return &ConductedInterview{}, nil
}
