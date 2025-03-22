// Maps database structures to protobuf structures

package api

import "interview/src/db"

func ConvertInterviewTemplateToProto(interviewTemplate *db.InterviewTemplate) *InterviewTemplate {
	return &InterviewTemplate{
		InterviewTemplateId: interviewTemplate.InterviewTemplateID[:],
		AverageScore:        interviewTemplate.AverageScore,
		AverageRating:       interviewTemplate.AverageRating,
		AmountConducted:     interviewTemplate.AmountConducted,
		Company:             interviewTemplate.Company,
		Role:                interviewTemplate.Role,
		Skills:              interviewTemplate.Skills,
		Description:         interviewTemplate.Description,
		Questions:           interviewTemplate.Questions,
		UserId:              interviewTemplate.UserID,
	}
}

func ConvertConductedInterviewToProto(conductedInterview *db.ConductedInterview) *ConductedInterview {
	return &ConductedInterview{
		ConductedInterviewId: conductedInterview.ConductedInterviewId[:],
		InterviewTemplateId:  conductedInterview.InterviewTemplateId[:],
		UserId:               conductedInterview.UserId,
		Score:                conductedInterview.Score,
		Rating:               conductedInterview.Rating,
		Role:                 conductedInterview.Role,
		Responses: &ResponseType{
			Feedback:  conductedInterview.Responses.Feedback,
			Answers:   conductedInterview.Responses.Answers,
			Questions: conductedInterview.Responses.Questions,
		},
	}
}
