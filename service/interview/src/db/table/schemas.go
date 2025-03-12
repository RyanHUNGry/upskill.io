package table

const INTERVIEW_TEMPLATES = `
CREATE TABLE IF NOT EXISTS interview_templates (
    interview_template_id TIMEUUID,
    average_score INT,
    average_rating INT,
    amount_conducted INT,
    company TEXT,
    role TEXT,
    skills SET<TEXT>,
    description TEXT,
    user_id INT,
    questions LIST<TEXT>,
    PRIMARY KEY (interview_template_id)
);
`

const AMOUNT_CONDUCTED_BY_INTERVIEW_TEMPLATE = `
CREATE TABLE IF NOT EXISTS amount_conducted_by_interview_template (
    interview_template_id TIMEUUID,
    amount_conducted COUNTER,
    PRIMARY KEY (interview_template_id)
);
`

const INTERVIEW_TEMPLATES_BY_COMPANY = `
CREATE TABLE IF NOT EXISTS interview_templates_by_company (
    company TEXT,
    interview_template_id TIMEUUID,
    PRIMARY KEY (company, interview_template_id)
) WITH CLUSTERING ORDER BY (interview_template_id DESC);
`

const INTERVIEW_TEMPLATES_BY_USER = `
CREATE TABLE IF NOT EXISTS interview_templates_by_user (
    user_id INT,
    interview_template_id TIMEUUID,
    PRIMARY KEY (user_id, interview_template_id)
) WITH CLUSTERING ORDER BY (interview_template_id DESC);
`

const AVERAGE_SCORES_BY_ROLE_AND_COMPANY = `
CREATE TABLE IF NOT EXISTS average_scores_by_role_and_company (
    role TEXT,
    company TEXT,
    average_score INT,
    interview_template_id TIMEUUID,
    PRIMARY KEY ((role, company), average_score, interview_template_id)
) WITH CLUSTERING ORDER BY (average_score DESC, interview_template_id DESC);
`

const AVERAGE_RATINGS_BY_ROLE_AND_COMPANY = `
CREATE TABLE IF NOT EXISTS average_ratings_by_role_and_company (
    role TEXT,
    company TEXT,
    average_rating INT,
    interview_template_id TIMEUUID,
    PRIMARY KEY ((role, company), average_rating, interview_template_id)
) WITH CLUSTERING ORDER BY (average_rating DESC, interview_template_id DESC);
`

const AMOUNT_CONDUCTED_BY_ROLE_AND_COMPANY = `
CREATE TABLE IF NOT EXISTS amount_conducted_by_role_and_company (
    role TEXT,
    company TEXT,
    amount_conducted INT,
	interview_template_id TIMEUUID,
    PRIMARY KEY ((role, company), amount_conducted, interview_template_id)
) WITH CLUSTERING ORDER BY (amount_conducted DESC, interview_template_id DESC);
`

const CONDUCTED_INTERVIEWS = `
CREATE TABLE IF NOT EXISTS conducted_interviews (
    conducted_interview_id TIMEUUID,
    interview_template_id TIMEUUID,
    score FLOAT,
    user_id INT,
    role TEXT,
    rating INT,
    responses frozen <response_type>,
    PRIMARY KEY (conducted_interview_id)
);
`
const RESPONSE_TYPE = `
CREATE TYPE IF NOT EXISTS response_type (
    questions LIST<TEXT>,
    answers LIST<TEXT>,
    feedback LIST<TEXT>
);
`

const CONDUCTED_INTERVIEWS_BY_USER = `
CREATE TABLE IF NOT EXISTS conducted_interviews_by_user (
    user_id INT,
    conducted_interview_id TIMEUUID,
    PRIMARY KEY (user_id, conducted_interview_id)
) WITH CLUSTERING ORDER BY (conducted_interview_id DESC);
`

var types = map[string]string{
	"response_type": RESPONSE_TYPE,
}

var schemas = map[string]string{
	"interview_templates":                    INTERVIEW_TEMPLATES,
	"interview_templates_by_company":         INTERVIEW_TEMPLATES_BY_COMPANY,
	"interview_templates_by_user":            INTERVIEW_TEMPLATES_BY_USER,
	"average_scores_by_role_and_company":     AVERAGE_SCORES_BY_ROLE_AND_COMPANY,
	"average_ratings_by_role_and_company":    AVERAGE_RATINGS_BY_ROLE_AND_COMPANY,
	"amount_conducted_by_role_and_company":   AMOUNT_CONDUCTED_BY_ROLE_AND_COMPANY,
	"conducted_interviews":                   CONDUCTED_INTERVIEWS,
	"conducted_interviews_by_user":           CONDUCTED_INTERVIEWS_BY_USER,
	"amount_conducted_by_interview_template": AMOUNT_CONDUCTED_BY_INTERVIEW_TEMPLATE,
}

var additionalCmds = map[string]string{
	"create_index_on_interview_templates": `
	CREATE INDEX IF NOT EXISTS ON interview_templates(role);
	`,
}
