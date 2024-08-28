package cassandra

import (
	"os"
)

// Depending on stage/cluster setup, alter replication settings
func (db *InterviewServiceDatabase) initializeKeyspace() {
	stage := os.Getenv("STAGE")

	switch stage {
	case "dev":
		queryString := "CREATE KEYSPACE IF NOT EXISTS interview_service WITH REPLICATION = {'class': 'NetworkTopologyStrategy', 'datacenter1': 3};"
		err := db.session.Query(queryString).WithContext(db.databaseContext).Exec()
		if err != nil {
			panic(err)
		}
	default:
		panic("Unknown stage")
	}
}

// Depending on stage/cluster setup, alter table settings
func (db *InterviewServiceDatabase) initializeTables() {
	for _, queryString := range schemas {
		err := db.session.Query(queryString).WithContext(db.databaseContext).Exec()
		if err != nil {
			panic(err)
		}
	}

	// implement seeding logic for test/dev env
}
