package api

import (
	context "context"
	"interview/src/db"
)

type InterviewServiceServerImpl struct {
	UnimplementedInterviewServiceServer
	session *db.CassandraSession
}

var mapper InterviewServiceMapper = InterviewServiceMapper{}

func (s *InterviewServiceServerImpl) CreateInterviewTemplateCall(ctx context.Context, in *CreateInterviewTemplate) (*InterviewTemplate, error) {
	interviewTemplateId := s.session.CreateInterviewTemplate(in.Company, in.Role, in.Skills, in.Description, in.Questions, in.UserId)
	interviewTemplate := mapper.ConvertInterviewTemplateToProto(s.session.QueryInterviewTemplate(interviewTemplateId))
	return interviewTemplate, nil
}

// func (s *InterviewServiceServerImpl) CreateConductedInterviewCall(ctx context.Context, in *CreateConductedInterview) (*ConductedInterview, error) {
// 	return &ConductedInterview{}, nil
// }
