package db

import (
	// do not import API package, it will cause a circular dependency
	// api package should reference this db package, but not the other way around
	// this ensures modularity and separation of concerns
	"interview/src/db/table"

	"github.com/gocql/gocql"
)

// creates interview template and updates associated tables
func (session *CassandraSession) CreateInterviewTemplate(
	company string,
	role string,
	skills []string,
	description string,
	questions []string,
	userId int32,
) gocql.UUID {
	interviewTemplateQuery := `
		INSERT INTO interview_templates
		(interview_template_id,
		average_score,
		average_rating,
		amount_conducted,
		company,
		role,
		skills,
		description,
		user_id,
		questions) VALUES
		(?, -1, -1, 0, ?, ?, ?, ?, ?, ?)
	`

	timeuuid := gocql.TimeUUID()

	err := session.Session.Query(
		interviewTemplateQuery,
		timeuuid,
		company,
		role,
		skills,
		description,
		userId,
		questions,
	).WithContext(session.Ctx).Exec()

	if err != nil {
		panic(err)
	}

	return timeuuid
}

func (session *CassandraSession) QueryInterviewTemplate(templateId gocql.UUID) map[string]interface{} {
	interviewTemplateQuery := `
		SELECT * FROM interview_templates WHERE interview_template_id = ?
	`

	resMap := map[string]interface{}{}

	for _, col := range table.InterviewTemplatesCols {
		res
	}

	err := session.Session.Query(interviewTemplateQuery, templateId).WithContext(session.Ctx).Scan()

	if err != nil {
		panic(err)
	}

	return &interviewTemplate
}
