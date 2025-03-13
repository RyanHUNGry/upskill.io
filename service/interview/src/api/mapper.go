package api

type InterviewServiceMapper struct{}

func (s InterviewServiceMapper) ConvertInterviewTemplateToProto(interviewTemplate map[string]interface{}) *InterviewTemplate {
	return &InterviewTemplate{
		InterviewTemplateId: interviewTemplate["interview_template_id"].([]byte),
		AverageScore:        interviewTemplate["average_score"].(int32),
		AverageRating:       interviewTemplate["average_rating"].(int32),
		AmountConducted:     interviewTemplate["amount_conducted"].(int32),
		Company:             interviewTemplate["company"].(string),
		Role:                interviewTemplate["role"].(string),
		Skills:              interviewTemplate["skills"].([]string),
		Description:         interviewTemplate["description"].(string),
		Questions:           interviewTemplate["questions"].([]string),
		UserId:              interviewTemplate["user_id"].(int32),
	}
}
