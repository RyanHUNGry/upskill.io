package cassandra

const (
	ratings = `
    CREATE TABLE IF NOT EXISTS ratings (
        interview_id UUID PRIMARY KEY,
		num_ratings counter,
        total_rating counter,
    );
    `
	interviews_by_user = `
    CREATE TABLE IF NOT EXISTS interviews_by_user (
        user_id UUID,
        interview_id UUID,
        company_name TEXT,
        rating FLOAT,
        PRIMARY KEY (user_id, interview_id)
    );
    `
	interviews_by_company = `
    CREATE TABLE IF NOT EXISTS interviews_by_company (
        company_name TEXT,
        interview_id UUID,
        rating FLOAT,
        user_id UUID,
        PRIMARY KEY (company_name, interview_id)
    );
    `
	top_interviews_by_rating = `
    CREATE TABLE IF NOT EXISTS top_interviews_by_rating (
        rating_bucket SMALLINT,
        actual_rating FLOAT,
        interview_id UUID,
        company_name TEXT,
        user_id UUID,
        PRIMARY KEY (rating, actual_rating, interview_id)
    );
    `
	questions_by_interview = `
    CREATE TABLE IF NOT EXISTS questions_by_interview (
        interview_id UUID,
        question_idx SMALLINT,
        question TEXT,
        company_name TEXT,
        PRIMARY KEY (interview_id, question_idx)
    );
    `
	answers_by_session = `
    CREATE TABLE IF NOT EXISTS answers_by_session (
        user_id UUID,
        session_id UUID,
        interview_id UUID,
        answer_idx SMALLINT,
        answer TEXT,
        question_idx SMALLINT,
        score FLOAT,
        PRIMARY KEY (user_id, session_id)
    );
    `
	sessions_by_user = `
    CREATE TABLE IF NOT EXISTS sessions_by_user (
        user_id UUID,
        session_id UUID,
        interview_id UUID,
        PRIMARY KEY (user_id, session_id)
    );
    `
)

var schemas = map[string]string{
	"ratings":                  ratings,
	"interviews_by_user":       interviews_by_user,
	"interviews_by_company":    interviews_by_company,
	"top_interviews_by_rating": top_interviews_by_rating,
	"questions_by_interview":   questions_by_interview,
	"answers_by_session":       answers_by_session,
	"sessisons_by_user":        sessions_by_user,
}
