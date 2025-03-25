package table

import "strings"

// Define schemas for tables as strings for GoCQL

// Programmatic way to define schemas for tables
func createTemplateQuery(tableName string, schema map[string]columnDetails) string {
	var query string = "CREATE TABLE IF NOT EXISTS " + tableName + " ("
	partitionKeys := make([]string, 0)
	clusteringKeys := make([]string, 0)
	clusteringKeyOrder := make([]string, 0)
	for column, columnType := range schema {
		pair := column + " " + columnType.columnType
		if columnType.isPartitionKey {
			partitionKeys = append(partitionKeys, column)
		} else if columnType.isClusteringKey {
			clusteringKeys = append(clusteringKeys, column)
			if columnType.isClusteringOrderDesc {
				clusteringKeyOrder = append(clusteringKeyOrder, column+" DESC")
			} else {
				clusteringKeyOrder = append(clusteringKeyOrder, column+" ASC")
			}
		}
		query += pair + ", "
	}

	primaryKey := ""
	if len(partitionKeys) > 1 {
		primaryKey += `(` + strings.Join(partitionKeys, ", ") + `)`
	} else {
		primaryKey += partitionKeys[0]
	}

	query += `PRIMARY KEY (` + primaryKey
	if len(clusteringKeys) > 0 {
		query += `, `
		query += strings.Join(clusteringKeys, ", ")
	}

	query += ")"

	if len(clusteringKeyOrder) > 0 {
		query += `) WITH CLUSTERING ORDER BY (` + strings.Join(clusteringKeyOrder, ", ")
	}

	query += `);`

	return query
}

// Define table schemas
var INTERVIEW_TEMPLATES string = createTemplateQuery("interview_templates", InterviewTemplatesColumns)
var INTERVIEW_TEMPLATES_BY_COMPANY string = createTemplateQuery("interview_templates_by_company", InterviewTemplatesByCompanyColumns)
var INTERVIEW_TEMPLATES_BY_USER string = createTemplateQuery("interview_templates_by_user", InterviewTemplatesByUserColumns)
var AVERAGE_SCORES_BY_ROLE_AND_COMPANY string = createTemplateQuery("average_scores_by_role_and_company", AverageScoresByRoleAndCompanyColumns)
var AVERAGE_RATINGS_BY_ROLE_AND_COMPANY string = createTemplateQuery("average_ratings_by_role_and_company", AverageRatingsByRoleAndCompanyColumns)
var AMOUNT_CONDUCTED_BY_ROLE_AND_COMPANY string = createTemplateQuery("amount_conducted_by_role_and_company", AmountConductedByRoleAndCompanyColumns)
var CONDUCTED_INTERVIEWS string = createTemplateQuery("conducted_interviews", ConductedInterviewsColumns)
var CONDUCTED_INTERVIEWS_BY_USER string = createTemplateQuery("conducted_interviews_by_user", ConductedInterviewsByUserColumns)

// Define UDTs
const RESPONSE_TYPE = `
CREATE TYPE IF NOT EXISTS response_type (
    questions LIST<TEXT>,
    answers LIST<TEXT>,
    feedback LIST<TEXT>
);
`

var Types = map[string]string{
	"response_type": RESPONSE_TYPE,
}

// Define additional table commands
var AdditionalCmds = map[string]string{
	"create_index_on_interview_templates": `
	CREATE INDEX IF NOT EXISTS ON interview_templates(role);
	`,
}

// Store schemas
var Schemas = map[string]string{
	"interview_templates":                  INTERVIEW_TEMPLATES,
	"interview_templates_by_company":       INTERVIEW_TEMPLATES_BY_COMPANY,
	"interview_templates_by_user":          INTERVIEW_TEMPLATES_BY_USER,
	"average_scores_by_role_and_company":   AVERAGE_SCORES_BY_ROLE_AND_COMPANY,
	"average_ratings_by_role_and_company":  AVERAGE_RATINGS_BY_ROLE_AND_COMPANY,
	"amount_conducted_by_role_and_company": AMOUNT_CONDUCTED_BY_ROLE_AND_COMPANY,
	"conducted_interviews":                 CONDUCTED_INTERVIEWS,
	"conducted_interviews_by_user":         CONDUCTED_INTERVIEWS_BY_USER,
}
