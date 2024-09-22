package cassandra

import (
	"context"
	"os"
	"time"

	"github.com/gocql/gocql"
)

type InterviewServiceDatabase struct {
	session         *gocql.Session
	databaseContext context.Context
}

func InitializeInterviewServiceDatabase(ctx context.Context) *InterviewServiceDatabase {
	stage := os.Getenv("STAGE")
	// Create keyspace initialization session
	clusterConfig := gocql.NewCluster("127.0.0.1")
	clusterConfig.Port = 9042
	clusterConfig.Timeout = time.Minute * 10
	keyspaceInitializationSession, err := clusterConfig.CreateSession()
	if err != nil {
		panic(err)
	}
	db := &InterviewServiceDatabase{session: keyspaceInitializationSession, databaseContext: ctx}

	err = db.session.Query("SELECT now() FROM system.local;").WithContext(ctx).Exec() // connection test query

	if err != nil {
		panic(err)
	}
	db.initializeKeyspace()
	db.CloseInterviewServiceDatabase()

	// Create usable session
	clusterConfig.Keyspace = "interview_service"

	if stage == "dev" {
		clusterConfig.Consistency = gocql.One // for development, use one because only one node
	} else {
		clusterConfig.Consistency = gocql.Quorum
	}

	session, err := clusterConfig.CreateSession()
	if err != nil {
		panic(err)
	}
	db = &InterviewServiceDatabase{session: session, databaseContext: ctx}
	err = db.session.Query("SELECT now() FROM system.local;").WithContext(ctx).Exec() // connection test query
	if err != nil {
		panic(err)
	}

	// Create tables
	db.initializeTables()
	return db
}

func (db *InterviewServiceDatabase) CloseInterviewServiceDatabase() {
	db.session.Close()
}
