package api

import (
	context "context"
	"interview/src/db"
	"interview/src/utils"

	"github.com/gocql/gocql"
)

type InterviewServiceServerImpl struct {
	UnimplementedInterviewServiceServer
	Database *db.Database
}

func (service *InterviewServiceServerImpl) CreateInterviewTemplateCall(ctx context.Context, in *CreateInterviewTemplate) (*InterviewTemplate, error) {
	interviewTemplateId, err := service.Database.CreateInterviewTemplate(in.Company, in.Role, in.Skills, in.Description, in.Questions, in.UserId)
	if err != nil {
		return nil, err
	}

	err = service.Database.InsertUserIdAndInterviewTemplateId(in.UserId, interviewTemplateId)
	if err != nil {
		return nil, err
	}

	err = service.Database.InsertRoleAndCompanyInterviewTemplateId(in.Role, in.Company, interviewTemplateId)
	if err != nil {
		return nil, err
	}

	interviewTemplate, _, err := service.Database.FindInterviewTemplateById(interviewTemplateId)
	if err != nil {
		return nil, err
	}

	return ConvertInterviewTemplateToProto(interviewTemplate), nil
}

func (service *InterviewServiceServerImpl) CreateConductedInterviewCall(ctx context.Context, in *CreateConductedInterview) (*ConductedInterview, error) {
	respType := db.ResponseType{Feedback: in.Responses.Feedback, Answers: in.Responses.Answers, Questions: in.Responses.Questions}

	conductedInterviewId, err := service.Database.CreateConductedInterview(in.InterviewTemplateId, in.UserId, in.Score, in.Rating, in.Role, in.Company, respType)
	if err != nil {
		return nil, err
	}

	averageScore, averageRating, amountConducted, err := service.Database.GetInterviewTemplateStats(in.InterviewTemplateId)
	if err != nil {
		return nil, err
	}

	newScore, newRating, newAmountConducted, err := service.Database.UpdateInterviewTemplateStats(in.InterviewTemplateId, averageScore, averageRating, amountConducted, in.Score, in.Rating)
	if err != nil {
		return nil, err
	}

	err = service.Database.UpdateRoleAndCompanyInterviewTemplateId(in.Role, in.Company, gocql.UUID(in.InterviewTemplateId), newScore, newRating, newAmountConducted)

	err = service.Database.InsertUserIdAndConductedInterviewId(in.UserId, conductedInterviewId)
	if err != nil {
		return nil, err
	}

	conductedInterview, _, err := service.Database.FindConductedInterviewById(conductedInterviewId)
	if err != nil {
		return nil, err
	}

	return ConvertConductedInterviewToProto(conductedInterview), nil
}

func (service *InterviewServiceServerImpl) GetConductedInterviewsByUserCall(ctx context.Context, in *GetConductedInterviewsByUser) (*ConductedInterviews, error) {
	conductedInterviewIds, err := service.Database.FindConductedInterviewIdsByUserId(in.UserId)
	if err != nil {
		return nil, err
	}

	_, conductedInterviews, err := service.Database.FindConductedInterviewById(conductedInterviewIds)

	if err != nil {
		return nil, err
	}

	return &ConductedInterviews{ConductedInterviews: utils.FunctionMap(conductedInterviews, ConvertConductedInterviewToProto)}, nil
}

func (service *InterviewServiceServerImpl) GetInterviewTemplatesByUserCall(ctx context.Context, in *GetInterviewTemplatesByUser) (*InterviewTemplates, error) {
	interviewTemplateIds, err := service.Database.FindInterviewTemplateIdsByUserId(in.UserId)
	if err != nil {
		return nil, err
	}

	_, interviewTemplates, err := service.Database.FindInterviewTemplateById(interviewTemplateIds)

	if err != nil {
		return nil, err
	}

	return &InterviewTemplates{InterviewTemplates: utils.FunctionMap(interviewTemplates, ConvertInterviewTemplateToProto)}, nil
}
