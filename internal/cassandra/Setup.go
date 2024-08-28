package cassandra

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
)

type InterviewServiceDatabase struct {
	session *gocql.Session
}

func InitializeInterviewServiceDatabase(ctx context.Context) *InterviewServiceDatabase {
	clusterConfig := gocql.NewCluster("127.0.0.1")
	clusterConfig.Port = 9042
	clusterConfig.Consistency = gocql.Quorum

	session, err := clusterConfig.CreateSession()
	if err != nil {
		panic(err)
	}

	db := &InterviewServiceDatabase{session: session}
	err = db.session.Query("SELECT now() FROM system.local;").WithContext(ctx).Exec() // connection test query
	if err != nil {
		panic(err)
	}
	fmt.Print("Database connection setup")
	return db
}

func (db *InterviewServiceDatabase) CloseInterviewServiceDatabase() {
	db.session.Close()
}
