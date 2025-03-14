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

func (service *InterviewServiceServerImpl) CreateConductedInterviewCall(ctx context.Context, in *CreateConductedInterview) (*ConductedInterview, error) {
	respType := db.ResponseType{Feedback: in.Responses.Feedback, Responses: in.Responses.Responses, Questions: in.Responses.Questions}
	conductedInterviewId, err := service.session.CreateConductedIntervew(in.InterviewTemplateId, in.UserId, in.Score, in.Rating, in.Role, respType)

	if err != nil {
		return nil, err
	}

	conductedInterview, err := service.session.FindConductedInterviewById(conductedInterviewId)
	if err != nil {
		return nil, err
	}

	return mapper.ConvertConductedInterviewToProto(conductedInterview), nil
}
