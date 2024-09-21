// mock data for the cassandra database
package cassandra

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gocql/gocql"
)

type sessionsByUserData struct {
	User_id      gocql.UUID
	Session_id   gocql.UUID
	Interview_id gocql.UUID
}

type answersBySessionData struct {
	User_id      gocql.UUID `fake:"-"` // this is a 16-byte or 128-bit UUID under the hood
	Session_id   gocql.UUID `fake:"-"`
	Interview_id gocql.UUID
	Answer_idx   int     `fake:"-"` // or use fake:"skip", same thing
	Answer       string  `fake:"{paragraph:5,4,12}"`
	Question_idx int     `fake:"-"`
	Score        float32 `fake:"{float32:60,100}"`
}

type interviewsByUserData struct {
	User_id      gocql.UUID `fake:"-"`
	Interview_id gocql.UUID `fake:"-"`
	Company_name string     `fake:"{company}"`
	Rating       float32    `fake:"{company}"`
}

type InterviewsByCompanyData struct {
	Company_name string
	Interview_id gocql.UUID
	Rating       float32
	User_id      gocql.UUID
}

// for use during initialization based on stage
func (db *InterviewServiceDatabase) seedDatabase() {
	gofakeit.Seed(69420) // ( ͡° ͜ʖ ͡°)
	/*
		sessions_by_user and answers_by_session creation
		Create 4 users, 3 sessions per user mapping to 3 different interviews, 3 answers per session/interview
	*/
	var sessionsByUserDataArray [4]*sessionsByUserData
	var answersBySessionDataArray [4][3]*answersBySessionData
	for i := 0; i < 4; i++ {
		sessionsByUserDataInstance := new(sessionsByUserData)
		err := gofakeit.Struct(sessionsByUserDataInstance)
		if err != nil {
			panic(err)
		}

		for j := 0; j < 3; j++ {
			answersBySessionDataInstance := new(answersBySessionData)
			err := gofakeit.Struct(answersBySessionDataInstance)
			if err != nil {
				panic(err)
			}

			answersBySessionDataInstance.User_id = sessionsByUserDataInstance.User_id
			answersBySessionDataInstance.Session_id = sessionsByUserDataInstance.Session_id
			answersBySessionDataInstance.Answer_idx = i + 1   // 1-indexed for ans/question idx
			answersBySessionDataInstance.Question_idx = i + 1 // 1-indexed for ans/question idx
			answersBySessionDataArray[i][j] = answersBySessionDataInstance
		}

		sessionsByUserDataArray[i] = sessionsByUserDataInstance
	}

	for _, sessions := range sessionsByUserDataArray {
		// IF NOT EXISTS requires Quorum
		insertionString := `INSERT INTO sessions_by_user (user_id, session_id, interview_id) VALUES (?, ?, ?) IF NOT EXISTS;` // IF NOT EXISTS does not alter the row if the primary key already exists. It will add a new row with that unique primary key that doesn't exist.
		err := db.session.Query(insertionString, sessions.User_id, sessions.Session_id, sessions.Interview_id).WithContext(db.databaseContext).Exec()
		if err != nil {
			panic(err)
		}
	}

	for _, answersArray := range answersBySessionDataArray {
		for _, answers := range answersArray {
			// IF NOT EXISTS requires Quorum
			insertionString := `INSERT INTO answers_by_session (user_id, session_id, interview_id, answer_idx, answer, question_idx, score) VALUES (?, ?, ?, ?, ?, ?, ?) IF NOT EXISTS;`
			err := db.session.Query(insertionString, answers.User_id, answers.Session_id, answers.Interview_id, answers.Answer_idx, answers.Answer, answers.Question_idx, answers.Score).WithContext(db.databaseContext).Exec()
			if err != nil {
				panic(err)
			}
		}
	}

	/*
		Create one new user with two interviews, each with three questions.
		Use the first user above to create two interviews, each with three questions.

		New user 1st interview -> 1st user in sessions_by_user
			- second interview not assigned to any user
		Original 1st user 1st interview -> 2nd user in sessions_by_user
		Original 1st user 2nd interview -> 3rd user in sessions_by_user
	*/

	fmt.Println("Database seeded ✅")
}
