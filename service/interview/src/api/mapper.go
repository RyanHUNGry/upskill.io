package api

import "interview/src/db"

type InterviewServiceMapper struct{}

func (s InterviewServiceMapper) ConvertInterviewTemplateToProto(interviewTemplate *db.InterviewTemplate) *InterviewTemplate {
	return &InterviewTemplate{
		InterviewTemplateId: interviewTemplate.InterviewTemplateID[:],
		AverageScore:        *interviewTemplate.AverageScore,
		AverageRating:       *interviewTemplate.AverageRating,
		AmountConducted:     *interviewTemplate.AmountConducted,
		Company:             *interviewTemplate.Company,
		Role:                *interviewTemplate.Role,
		Skills:              interviewTemplate.Skills,
		Description:         *interviewTemplate.Description,
		Questions:           interviewTemplate.Questions,
		UserId:              *interviewTemplate.UserID,
	}
}
