// CassandraDBconnection session creation
package db

import (
	"context"
	"fmt"
	"interview/src/db/table"
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

type Database struct {
	Session *gocql.Session
	Ctx     context.Context
}

// Initializes Cassandra session, creating interview keyspace if it doesn't exist
func Connect(host string, port string, ctx context.Context) (*Database, error) {
	cluster := gocql.NewCluster(host + ":" + port)
	cluster.ConnectTimeout = 3 * time.Second
	cluster.Logger = log.New(os.Stdout, "gocql: ", log.LstdFlags)
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}

	err = session.Query("SELECT uuid() FROM system.local;").WithContext(ctx).Exec()

	if err != nil {
		log.Fatalf("Failed database healthcheck: %v", err)
	}

	// Create interview keyspace if necessary
	var keyspace string
	err = session.Query("SELECT keyspace_name FROM system_schema.keyspaces WHERE keyspace_name = 'interview';").WithContext(ctx).Scan(&keyspace)

	if err != nil && err.Error() != "not found" {
		log.Fatalf("Failed to check if keyspace exists: %v", err)
	}

	keyspaceExists := keyspace == "interview"
	if !keyspaceExists {
		err := session.Query(`CREATE KEYSPACE interview 
			WITH REPLICATION = { 
				'class': 'SimpleStrategy', 
				'replication_factor': 1 
			}`).WithContext(ctx).Exec()

		if err != nil {
			log.Fatalf("Failed to create keyspace: %v", err)
		}

		session.Close()
	}

	cluster.Keyspace = "interview"
	session, err = cluster.CreateSession()

	if err != nil {
		log.Fatal("Failed database connection")
	}

	return &Database{Session: session, Ctx: ctx}, nil
}

// Initialize any tables (and any secondary indexes, types, roles, materialized views, and etc...) not present in database based on schemas
func (db *Database) InitializeTables() error {
	// Initialize any UDTs
	scanner := db.Session.Query("SELECT type_name FROM system_schema.types WHERE keyspace_name = 'interview';").WithContext(db.Ctx).Iter().Scanner()
	typeNames := map[string]struct{}{}
	for scanner.Next() {
		var typeName string
		err := scanner.Scan(&typeName)
		if err != nil {
			return err
		}
		typeNames[typeName] = struct{}{}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	for typeName := range table.Types {
		if _, ok := typeNames[typeName]; !ok {
			fmt.Printf("Creating type %s\n", typeName)
			if err := db.Session.Query(table.Types[typeName]).Exec(); err != nil {
				return err
			}
		}
	}

	// Initialize any tables
	scanner = db.Session.Query("SELECT table_name FROM system_schema.tables WHERE keyspace_name = 'interview';").WithContext(db.Ctx).Iter().Scanner()
	tableNames := map[string]struct{}{}
	for scanner.Next() {
		var tableName string
		err := scanner.Scan(&tableName)
		if err != nil {
			return err
		}
		tableNames[tableName] = struct{}{}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	for tableName := range table.Schemas {
		if _, ok := tableNames[tableName]; !ok {
			fmt.Printf("Creating table %s\n", tableName)
			if err := db.Session.Query(table.Schemas[tableName]).Exec(); err != nil {
				return err
			}
		}
	}

	// Execute any additional commands
	for cmdTitle, cmd := range table.AdditionalCmds {
		if err := db.Session.Query(cmd).Exec(); err != nil {
			return err
		} else {
			fmt.Println("Command", cmdTitle, "executed successfully")
		}
	}

	return nil
}

// Drop all tables, but retains any UDTs
func (db *Database) DropAllTables() error {
	fmt.Println("DROPPING ALL TABLES...")

	tableNames := []string{}
	scanner := db.Session.Query(`SELECT table_name FROM system_schema.tables WHERE keyspace_name='interview';`).WithContext(db.Ctx).Iter().Scanner()
	for scanner.Next() {
		var tableName string
		err := scanner.Scan(&tableName)
		if err != nil {
			log.Fatal("Failed to scan table name", err)
		}

		tableNames = append(tableNames, tableName)
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	for _, tableName := range tableNames {
		err := db.Session.Query(`DROP TABLE ` + tableName).WithContext(db.Ctx).Exec()
		if err != nil {
			return err
		}
	}

	return nil
}
