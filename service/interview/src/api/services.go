package api

import (
	context "context"
	"interview/src/db"
)

type InterviewServiceServerImpl struct {
	UnimplementedInterviewServiceServer
	session *db.Database
}

var mapper InterviewServiceMapper = InterviewServiceMapper{}

func (service *InterviewServiceServerImpl) CreateInterviewTemplateCall(ctx context.Context, in *CreateInterviewTemplate) (*InterviewTemplate, error) {
	interviewTemplateId, err := service.session.CreateInterviewTemplate(in.Company, in.Role, in.Skills, in.Description, in.Questions, in.UserId)

	if err != nil {
		return nil, err
	}

	interviewTemplate, err := service.session.FindInterviewTemplateById(interviewTemplateId)

	if err != nil {
		return nil, err
	}

	return mapper.ConvertInterviewTemplateToProto(interviewTemplate), nil
}

// func (s *InterviewServiceServerImpl) CreateConductedInterviewCall(ctx context.Context, in *CreateConductedInterview) (*ConductedInterview, error) {
// 	return &ConductedInterview{}, nil
// }
