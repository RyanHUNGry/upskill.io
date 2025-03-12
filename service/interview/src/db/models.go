package db

import (
	"fmt"
	"strings"
)

// creates interview template and updates associated tables
func (session *CassandraSession) CreateInterviewTemplate(
	company string,
	role string,
	skills []string,
	description string,
	questions []string,
	userId int32,
) {
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
		(timeuuid(), -1, -1, 0, ?, ?, ?, ?, ?, ?)
	`

	commaString := strings.Join(skills, ",")
	commaString = "{ " + commaString + " }"
	fmt.Println(commaString)

	err := session.Session.Query(
		interviewTemplateQuery,
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
}
