package db

import (
	// do not import API package, it will cause a circular dependency
	// api package should reference this db package, but not the other way around
	// this ensures modularity and separation of concerns

	"github.com/gocql/gocql"
)

type InterviewTemplate struct {
	InterviewTemplateID gocql.UUID `json:"interview_template_id"`
	AverageScore        *int32     `json:"average_score"`
	AverageRating       *int32     `json:"average_rating"`
	AmountConducted     *int32     `json:"amount_conducted"`
	Company             *string    `json:"company"`
	Role                *string    `json:"role"`
	Skills              []string   `json:"skills"`
	Description         *string    `json:"description"`
	UserID              *int32     `json:"user_id"`
	Questions           []string   `json:"questions"`
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
