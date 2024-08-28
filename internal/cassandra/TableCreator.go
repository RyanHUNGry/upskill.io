package cassandra

import "fmt"

func (db *InterviewServiceDatabase) InitializeTables() {
	// use an env var to see if we want to seed, which is good in dev/test envs
	fmt.Print("Populating tables")
}
