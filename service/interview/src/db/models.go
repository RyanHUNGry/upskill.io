package db

import (
	// do not import API package, it will cause a circular dependency
	// api package should reference this db package, but not the other way around
	// this ensures modularity and separation of concerns

	"github.com/gocql/gocql"
)

type InterviewTemplate struct {
	InterviewTemplateID gocql.UUID
	AverageScore        int32
	AverageRating       int32
	AmountConducted     int32
	Company             string
	Role                string
	Skills              []string
	Description         string
	UserID              int32
	Questions           []string
}

// using struct (un)marshalling
type ResponseType struct {
	Feedback  []string `cql:"feedback"`
	Answers   []string `cql:"answers"`
	Questions []string `cql:"questions"`
}

type ConductedInterview struct {
	ConductedInterviewId gocql.UUID
	InterviewTemplateId  gocql.UUID
	UserId               int32
	Score                int32
	Rating               int32
	Role                 string
	Responses            ResponseType
}

// creates interview template and updates associated tables
func (db *Database) CreateInterviewTemplate(
	company string,
	role string,
	skills []string,
	description string,
	questions []string,
	userId int32,
) (gocql.UUID, error) {
	query := `
	INSERT INTO interview_templates (
		interview_template_id,
		average_score,
		average_rating,
		amount_conducted,
		company,
		role,
		skills,
		description,
		user_id,
		questions
	) VALUES (
		?, -1, -1, 0, ?, ?, ?, ?, ?, ?
	)
	`

	timeuuid := gocql.TimeUUID()

	err := db.Session.Query(
		query,
		timeuuid,
		company,
		role,
		skills,
		description,
		userId,
		questions,
	).WithContext(db.Ctx).Exec()

	if err != nil {
		return gocql.UUID{}, err
	}

	return timeuuid, nil
}

func (db *Database) FindInterviewTemplateById(templateId gocql.UUID) (*InterviewTemplate, error) {
	query := `SELECT * FROM interview_templates WHERE interview_template_id = ?`

	var template InterviewTemplate

	err := db.Session.Query(query, templateId).WithContext(db.Ctx).Scan(
		&template.InterviewTemplateID,
		&template.AmountConducted,
		&template.AverageRating,
		&template.AverageScore,
		&template.Company,
		&template.Description,
		&template.Questions,
		&template.Role,
		&template.Skills,
		&template.UserID,
	)

	if err != nil {
		return nil, err
	}

	return &template, nil
}

func (db *Database) CreateConductedIntervew(
	interviewTemplateId []byte,
	userId int32,
	score int32,
	rating int32,
	role string,
	responses ResponseType,
) (gocql.UUID, error) {
	query := `
	INSERT INTO conducted_interviews (
		interview_template_id,
		conducted_interview_id,
		score,
		rating,
		role,
		user_id,
		responses
	) VALUES (
		?, ?, ?, ?, ?, ?, ?
	)
	`

	conductedInterviewId := gocql.TimeUUID()
	err := db.Session.Query(
		query,
		interviewTemplateId,
		conductedInterviewId,
		score,
		rating,
		role,
		userId,
		responses,
	).WithContext(db.Ctx).Exec()

	if err != nil {
		return gocql.UUID{}, err
	}

	return conductedInterviewId, nil
}

func (db *Database) FindConductedInterviewById(conductedInterviewId gocql.UUID) (*ConductedInterview, error) {
	order := `conducted_interview_id,
	interview_template_id,
	score,
	user_id,
	role,
	rating,
	responses`
	query := `SELECT ` + order + ` FROM conducted_interviews WHERE conducted_interview_id = ?`

	var conductedInterview ConductedInterview

	err := db.Session.Query(query, conductedInterviewId).WithContext(db.Ctx).Scan(
		&conductedInterview.ConductedInterviewId,
		&conductedInterview.InterviewTemplateId,
		&conductedInterview.Score,
		&conductedInterview.UserId,
		&conductedInterview.Role,
		&conductedInterview.Rating,
		&conductedInterview.Responses,
	)

	if err != nil {
		return nil, err
	}

	return &conductedInterview, nil
}
