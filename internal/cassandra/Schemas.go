package cassandra

const (
	getUsersByStateSchema = `
    CREATE TABLE IF NOT EXISTS get_users_by_state (
        state UUID PRIMARY KEY,
		uid UUID,
        name TEXT,
        email TEXT
    );
    `
)

var schemas = map[string]string{
	"get_users_by_state": getUsersByStateSchema,
}

// Queries.
// get interview sessions by user, so partition on user ID. Also order by date created
//

// get all resume details by user
//
// get interview sessions by user
// get prompt details/company by interview session
// get
