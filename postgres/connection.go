package postgres

import (
	"database/sql"
	"fmt"

	"sensei/utilities"

	_ "github.com/lib/pq"
)

// Struct containing all the required details to connect to
// a Postgres instance
type PgConnectionDetails struct {
	// Host of the postgres instance to connect to
	host string

	// Port at which the postgres instance is running at
	port int

	// User with which to connect to the instance
	user string

	// Password for the user supplied above
	password string
}

// Generates a connection to a Postgres instance, returns with a pointer
// to the connection instance.
func PgConnect(options PgConnectionDetails) *sql.DB {
	var err error

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=sensei sslmode=disable",
		options.host,
		options.port,
		options.user,
		options.password,
	)

	db, err := sql.Open("postgres", connStr)
	utilities.PanicWithError(err)

	err = db.Ping()
	utilities.PanicWithError(err)

	return db
}
