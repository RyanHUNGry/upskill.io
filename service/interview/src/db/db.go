// CassandraDBconnection session creation
package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

type Database struct {
	Session *gocql.Session
	Ctx     context.Context
}

// Initializes cassandra session, creates interview keyspace if not exist
func Connect(host string, port string, ctx context.Context) (*Database, error) {
	cluster := gocql.NewCluster(host + ":" + port)
	cluster.ConnectTimeout = 3 * time.Second
	cluster.Logger = log.New(os.Stdout, "gocql: ", log.LstdFlags) // output to stdout, prefix with gocql:, add timestamp
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal("Failed database connection")
	}

	err = session.Query("SELECT uuid() FROM system.local;").WithContext(ctx).Exec()

	if err != nil {
		log.Fatal("Failed database healthcheck")
	}

	var keyspace string
	err = session.Query("SELECT keyspace_name FROM system_schema.keyspaces WHERE keyspace_name = 'interview';").WithContext(ctx).Scan(&keyspace)

	if err != nil && err.Error() != "not found" {
		log.Fatal("Failed to check keyspace existence: ", err)
	}

	keyspaceExists := keyspace == "interview"

	if !keyspaceExists {
		err := session.Query(`CREATE KEYSPACE interview 
			WITH REPLICATION = { 
				'class': 'SimpleStrategy', 
				'replication_factor': 1 
			}`).WithContext(ctx).Exec()

		if err != nil {
			log.Fatal("Failed to create keyspace: ", err)
		}

		session.Close() //close the current session as this was the initial session to create keyspace
	}

	cluster.Keyspace = "interview"
	session, err = cluster.CreateSession()

	if err != nil {
		log.Fatal("Failed database connection")
	}

	return &Database{Session: session, Ctx: ctx}, nil
}
