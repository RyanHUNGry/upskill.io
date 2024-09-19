package cassandra

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gocql/gocql"
)

type sessionsByUserData struct {
	User_id    gocql.UUID
	Session_id gocql.UUID
}

type answersBySessionData struct {
	User_id       gocql.UUID `fake:"-"` // this is a 16-byte or 128-bit UUID under the hood
	Session_id    gocql.UUID `fake:"-"`
	Interfview_id gocql.UUID
	Answer_idx    int     `fake:"-"` // or use fake:"skip", same thing
	Answer        string  `fake:"{paragraph:5,4,12}"`
	Question_idx  int     `fake:"-"`
	Score         float32 `fake:"{float32:60,100}"`
}

// for use during initialization based on stage
func (db *InterviewServiceDatabase) seedDatabase() {
	gofakeit.Seed(69420) // ( ͡° ͜ʖ ͡°)
	/*
		Create 4 users, 3 sessions per user, 3 answers per session
	*/
	var sessionsByUserDataArray [4]sessionsByUserData
	var answersBySessionDataArray
	for i := 0; i < 4; i++ {
		sessionsByUserDataInstance := new(sessionsByUserData)
		err := gofakeit.Struct(sessionsByUserDataInstance)
		if err != nil {
			panic(err)
		}

		for i := 0; i < 3; i++ {
			answersBySessionDataInstance := new(answersBySessionData)
			err := gofakeit.Struct(answersBySessionDataInstance)
			if err != nil {
				panic(err)
			}

			answersBySessionDataInstance.User_id = sessionsByUserDataInstance.User_id
			answersBySessionDataInstance.Session_id = sessionsByUserDataInstance.Session_id
			answersBySessionDataInstance.Answer_idx = i + 1 // 1-indexed for ans/question idx
			answersBySessionDataInstance.Question_idx = i + 1
		}
	}

	err := gofakeit.Struct(answersBySessionDataInstance)

	if err != nil {
		panic(err)
	}

	err = gofakeit.Struct(sessionsByUserDataInstance)
	if err != nil {
		panic(err)
	}

	fmt.Println(answersBySessionDataInstance.User_id, answersBySessionDataInstance.Score)

	fmt.Println("Database seeded ✅")
}
