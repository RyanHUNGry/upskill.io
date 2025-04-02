package db

import (
	// do not import API package, it will cause a circular dependency
	// api package should reference this db package, but not the other way around
	// this ensures modularity and separation of concerns

	"errors"
	"interview/src/utils"
	"strings"

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
	Company              string
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
	columns := []string{
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

	conditions := make([]string, 0, len(columns))
	for range columns {
		conditions = append(conditions, "?")
	}

	query := `INSERT INTO interview_templates (` + strings.Join(columns, ", ") + `) VALUES (` + strings.Join(conditions, ", ") + `)`

	interviewTemplateId := gocql.TimeUUID()
	averageScore := -1
	averageRating := -1
	amountConducted := 0

	err := db.Session.Query(
		query,
		interviewTemplateId,
		averageScore,
		averageRating,
		amountConducted,
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

	return interviewTemplateId, nil
}

// Find single interview template by ID, or find a set of interview templates by multiple IDs
func (db *Database) FindInterviewTemplateById(interviewTemplateId any) (*InterviewTemplate, []*InterviewTemplate, error) {
	columns := []string{
		"interview_template_id",
		"amount_conducted",
		"average_rating",
		"average_score",
		"company",
		"description",
		"questions",
		"role",
		"skills",
		"user_id",
	}

	// Convert byte slice to UUID type
	if v, ok := interviewTemplateId.([]byte); ok {
		interviewTemplateId = gocql.UUID(v)
	}

	if _, ok := interviewTemplateId.(gocql.UUID); ok {
		interviewTemplateId := interviewTemplateId.(gocql.UUID)
		query := `SELECT ` + strings.Join(columns, ", ") + ` FROM interview_templates WHERE interview_template_id = ?`
		var interviewTemplate InterviewTemplate
		err := db.Session.Query(query, interviewTemplateId).WithContext(db.Ctx).Scan(
			&interviewTemplate.InterviewTemplateID,
			&interviewTemplate.AmountConducted,
			&interviewTemplate.AverageRating,
			&interviewTemplate.AverageScore,
			&interviewTemplate.Company,
			&interviewTemplate.Description,
			&interviewTemplate.Questions,
			&interviewTemplate.Role,
			&interviewTemplate.Skills,
			&interviewTemplate.UserID,
		)

		if err != nil {
			return nil, nil, err
		}

		return &interviewTemplate, nil, nil
	} else if interviewTemplateIds, ok := interviewTemplateId.([]gocql.UUID); ok {
		conditions := make([]string, 0, len(interviewTemplateIds))
		for range interviewTemplateIds {
			conditions = append(conditions, "?")
		}

		query := `SELECT ` + strings.Join(columns, ", ") + ` FROM interview_templates WHERE interview_template_id in (` + strings.Join(conditions, ", ") + `);`

		// Under the hood, the variadic arguments for Query() is a []interface{} so cast IDs
		scanner := db.Session.Query(query, utils.AnySliceConverter(interviewTemplateIds)...).WithContext(db.Ctx).Iter().Scanner()
		var interviewTemplates []*InterviewTemplate

		for scanner.Next() {
			var interviewTemplate InterviewTemplate
			err := scanner.Scan(
				&interviewTemplate.InterviewTemplateID,
				&interviewTemplate.AmountConducted,
				&interviewTemplate.AverageRating,
				&interviewTemplate.AverageScore,
				&interviewTemplate.Company,
				&interviewTemplate.Description,
				&interviewTemplate.Questions,
				&interviewTemplate.Role,
				&interviewTemplate.Skills,
				&interviewTemplate.UserID,
			)

			interviewTemplates = append(interviewTemplates, &interviewTemplate)

			if err != nil {
				return nil, nil, err
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, nil, err
		}

		return nil, interviewTemplates, nil
	} else {
		return nil, nil, errors.New("invalid type for interviewTemplateId")
	}
}

func (db *Database) CreateConductedInterview(
	interviewTemplateId []byte,
	userId int32,
	score int32,
	rating int32,
	role string,
	company string,
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

// Find single conducted interview by ID, or find a set of conducted interview by multiple IDs
func (db *Database) FindConductedInterviewById(conductedInterviewId any) (*ConductedInterview, []*ConductedInterview, error) {
	columns := []string{
		"conducted_interview_id",
		"interview_template_id",
		"score",
		"user_id",
		"role",
		"rating",
		"responses",
	}

	if _, ok := conductedInterviewId.(gocql.UUID); ok {
		conductedInterviewId := conductedInterviewId.(gocql.UUID)
		query := `SELECT ` + strings.Join(columns, ", ") + ` FROM conducted_interviews WHERE conducted_interview_id = ?`
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
			return nil, nil, err
		}

		return &conductedInterview, nil, err
	} else if conductedInterviewIds, ok := conductedInterviewId.([]gocql.UUID); ok {
		conditions := make([]string, 0, len(conductedInterviewIds))
		for range conductedInterviewIds {
			conditions = append(conditions, "?")
		}

		query := `SELECT ` + strings.Join(columns, ", ") + ` FROM conducted_interviews WHERE conducted_interview_id in (` + strings.Join(conditions, ", ") + `);`

		// Under the hood, the variadic arguments for Query() is a []interface{} so cast IDs
		scanner := db.Session.Query(query, utils.AnySliceConverter(conductedInterviewIds)...).WithContext(db.Ctx).Iter().Scanner()
		var conductedInterviews []*ConductedInterview

		for scanner.Next() {
			var conductedInterview ConductedInterview
			err := scanner.Scan(
				&conductedInterview.ConductedInterviewId,
				&conductedInterview.InterviewTemplateId,
				&conductedInterview.Score,
				&conductedInterview.UserId,
				&conductedInterview.Role,
				&conductedInterview.Rating,
				&conductedInterview.Responses,
			)

			conductedInterviews = append(conductedInterviews, &conductedInterview)

			if err != nil {
				return nil, nil, err
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, nil, err
		}

		return nil, conductedInterviews, nil
	} else {
		return nil, nil, errors.New("invalid type for conductedInterviewId")
	}
}

func (db *Database) FindInterviewTemplateIdsByUserId(userId int32) (interviewTemplateIds []gocql.UUID, err error) {
	columns := []string{
		"user_id",
		"interview_template_id",
	}
	query := `SELECT ` + strings.Join(columns, ", ") + ` FROM interview_templates_by_user WHERE user_id = ?`
	scanner := db.Session.Query(query, userId).WithContext(db.Ctx).Iter().Scanner()

	for scanner.Next() {
		var user_id int32
		var interview_template_id gocql.UUID

		err := scanner.Scan(&user_id, &interview_template_id)

		interviewTemplateIds = append(interviewTemplateIds, interview_template_id)
		if err != nil {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return interviewTemplateIds, nil
}

func (db *Database) FindConductedInterviewIdsByUserId(userId int32) (conductedInterviewIds []gocql.UUID, err error) {
	columns := []string{
		"user_id",
		"conducted_interview_id",
	}
	query := `SELECT ` + strings.Join(columns, ", ") + ` FROM conducted_interviews_by_user WHERE user_id = ?`
	scanner := db.Session.Query(query, userId).WithContext(db.Ctx).Iter().Scanner()

	for scanner.Next() {
		var user_id int32
		var conducted_interview_id gocql.UUID

		err := scanner.Scan(&user_id, &conducted_interview_id)

		conductedInterviewIds = append(conductedInterviewIds, conducted_interview_id)
		if err != nil {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return conductedInterviewIds, nil
}

func (db *Database) InsertUserIdAndConductedInterviewId(userId int32, conductedInterviewId gocql.UUID) error {
	columns := []string{"conducted_interview_id", "user_id"}
	query := `INSERT INTO conducted_interviews_by_user (` + strings.Join(columns, ", ") + `) VALUES (?, ?)`

	err := db.Session.Query(query, conductedInterviewId, userId).WithContext(db.Ctx).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) InsertUserIdAndInterviewTemplateId(userId int32, interviewTemplateId gocql.UUID) error {
	columns := []string{"interview_template_id", "user_id"}
	query := `INSERT INTO interview_templates_by_user (` + strings.Join(columns, ", ") + `) VALUES (?, ?)`

	err := db.Session.Query(query, interviewTemplateId, userId).WithContext(db.Ctx).Exec()
	if err != nil {
		return err
	}

	return nil
}

// Inserts the role and company filtering tables
func (db *Database) InsertRoleAndCompanyInterviewTemplateId(role string, company string, interviewTemplateId gocql.UUID) error {
	columns := []string{"role", "company", "interview_template_id", "average_score"}
	query := `INSERT INTO average_scores_by_role_and_company (` + strings.Join(columns, ", ") + `) VALUES ` + utils.GenerateConditions(columns)
	err := db.Session.Query(
		query, role, company, interviewTemplateId, -1,
	).WithContext(db.Ctx).Exec()
	if err != nil {
		return err
	}

	columns = []string{"role", "company", "interview_template_id", "average_rating"}
	query = `INSERT INTO average_ratings_by_role_and_company (` + strings.Join(columns, ", ") + `) VALUES ` + utils.GenerateConditions(columns)
	err = db.Session.Query(
		query, role, company, interviewTemplateId, -1,
	).WithContext(db.Ctx).Exec()
	if err != nil {
		return err
	}

	columns = []string{"role", "company", "interview_template_id", "amount_conducted"}
	query = `INSERT INTO amount_conducted_by_role_and_company (` + strings.Join(columns, ", ") + `) VALUES ` + utils.GenerateConditions(columns)
	err = db.Session.Query(
		query, role, company, interviewTemplateId, 0,
	).WithContext(db.Ctx).Exec()
	if err != nil {
		return err
	}

	return nil
}

// Target tables: average_scores_by_role_and_company, average_ratings_by_role_and_company, amount_conducted_by_role_and_company
// Updates the role and company filtering tables
func (db *Database) UpdateRoleAndCompanyInterviewTemplateId(role string, company string, interviewTemplateId gocql.UUID, averageScore int32, averageRating int32, amountConducted int32) error {
	query := `UPDATE average_scores_by_role_and_company SET average_score = ? WHERE role = ? AND company = ? AND interview_template_id = ?`
	err := db.Session.Query(
		query, averageScore, role, company, interviewTemplateId,
	).WithContext(db.Ctx).Exec()
	if err != nil {
		return err
	}

	query = `UPDATE average_ratings_by_role_and_company SET average_rating = ? WHERE role = ? AND company = ? AND interview_template_id = ?`
	err = db.Session.Query(
		query, averageRating, role, company, interviewTemplateId,
	).WithContext(db.Ctx).Exec()
	if err != nil {
		return err
	}

	query = `UPDATE amount_conducted_by_role_and_company SET amount_conducted = ? WHERE role = ? AND company = ? AND interview_template_id = ?`
	err = db.Session.Query(
		query, amountConducted, role, company, interviewTemplateId,
	).WithContext(db.Ctx).Exec()
	if err != nil {
		return err
	}

	return nil
}

// Target tables: interview_templates
// Get interview template statistics
func (db *Database) GetInterviewTemplateStats(interviewTemplateId []byte) (averageScore int32, averageRating int32, amountConducted int32, err error) {
	columns := []string{"average_score", "average_rating", "amount_conducted"}
	query := `SELECT ` + strings.Join(columns, ", ") + ` FROM interview_templates WHERE interview_template_id = ?`

	err = db.Session.Query(query, interviewTemplateId).WithContext(db.Ctx).Scan(
		&averageScore,
		&averageRating,
		&amountConducted,
	)
	if err != nil {
		return -1, -1, -1, err
	}

	return averageScore, averageRating, amountConducted, nil
}

// Target tables: interview_templates
// Update interview template statistics given a single incoming score and rating and the existing statistics
func (db *Database) UpdateInterviewTemplateStats(interviewTemplateId []byte, averageScore int32, averageRating int32, amountConducted int32, score int32, rating int32) (int32, int32, int32, error) {
	if averageScore == -1 || averageRating == -1 {
		averageScore = 0
		averageRating = 0
	}

	newScore := (averageScore*amountConducted + score) / (amountConducted + 1)
	newRating := (averageRating*amountConducted + rating) / (amountConducted + 1)
	newAmountConducted := amountConducted + 1

	query := `UPDATE interview_templates SET average_score = ?, average_rating = ?, amount_conducted = ? WHERE interview_template_id = ?`
	err := db.Session.Query(query, newScore, newRating, newAmountConducted, interviewTemplateId).WithContext(db.Ctx).Exec()
	if err != nil {
		return -1, -1, -1, err
	}

	return newScore, newRating, newAmountConducted, nil
}
