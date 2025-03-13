package table

var InterviewTemplatesCols = []string{
	"interview_template_id",
	"average_score",
	"average_rating",
	"amount_conducted",
	"company",
	"role",
	"skills",
	"description",
	"user_id",
	"questions",
}

var AmountConductedByInterviewTemplateCols = []string{
	"interview_template_id",
	"amount_conducted",
}

var InterviewTemplatesByCompanyCols = []string{
	"company",
	"interview_template_id",
}

var InterviewTemplatesByUserCols = []string{
	"user_id",
	"interview_template_id",
}

var AverageScoresByRoleAndCompanyCols = []string{
	"role",
	"company",
	"average_score",
	"interview_template_id",
}

var AverageRatingsByRoleAndCompanyCols = []string{
	"role",
	"company",
	"average_rating",
	"interview_template_id",
}

var AmountConductedByRoleAndCompanyCols = []string{
	"role",
	"company",
	"amount_conducted",
	"interview_template_id",
}

var ConductedInterviewsCols = []string{
	"conducted_interview_id",
	"interview_template_id",
	"score",
	"user_id",
	"role",
	"rating",
	"responses",
}

var ConductedInterviewsByUserCols = []string{
	"user_id",
	"conducted_interview_id",
}
