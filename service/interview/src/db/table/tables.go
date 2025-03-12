package table

import (
	"context"
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func InitializeTables(session *gocql.Session, ctx context.Context) {

	var table string
	err := session.Query("SELECT table_name FROM system_schema.tables WHERE keyspace_name = 'interview';").WithContext(ctx).Scan(&table)
	tablesCreated := table != ""

	if err != nil && err.Error() != "not found" {
		log.Fatal("Failed to check if tables exist", err)
	}

	if tablesCreated {
		return // don't run CREATE statements
	}

	for customType, schema := range types {
		if err := session.Query(schema).Exec(); err != nil {
			log.Fatal("Failed to create custom type", customType, err)
		} else {
			fmt.Println("Custom type", customType, "created successfully")
		}
	}

	for tableTitle, schema := range schemas {
		if err := session.Query(schema).Exec(); err != nil {
			log.Fatal("Failed to create table", tableTitle, err)
		} else {
			fmt.Println("Table", tableTitle, "created successfully")
		}
	}

	for cmdTitle, cmd := range additionalCmds {
		if err := session.Query(cmd).Exec(); err != nil {
			log.Fatal("Failed to execute command", cmdTitle, err)
		} else {
			fmt.Println("Command", cmdTitle, "executed successfully")
		}
	}
}

func DropAllTables(session *gocql.Session, ctx context.Context) {
	getTableNamesQuery := `SELECT table_name FROM system_schema.tables WHERE keyspace_name='interview';`
	tableNames := []string{}

	scanner := session.Query(getTableNamesQuery).WithContext(ctx).Iter().Scanner()
	for scanner.Next() {
		var tableName string
		err := scanner.Scan(&tableName)
		if err != nil {
			log.Fatal("Failed to scan table name", err)
		}

		tableNames = append(tableNames, tableName)
	}

	if scanner.Err() != nil {
		log.Fatal("Failed to scan table names", scanner.Err())
	}

	for _, tableName := range tableNames {
		dropTableQuery := fmt.Sprintf("DROP TABLE %s;", tableName)
		err := session.Query(dropTableQuery).WithContext(ctx).Exec()
		if err != nil {
			log.Fatal("Failed to drop table", tableName, err)
		}
	}
}
