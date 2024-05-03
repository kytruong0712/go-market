package pg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	pkgerrors "github.com/pkg/errors"
)

// DatabaseName indicates name of applying database. It's `postgres` in this case
const DatabaseName = "postgres"

// Connect connects to database
func Connect(dbURL string) (*sql.DB, error) {
	conn, err := sql.Open(DatabaseName, dbURL)
	if err != nil {
		os.Exit(1)

		return nil, pkgerrors.WithStack(fmt.Errorf("connect DB failed. Err: %w", err))
	}

	if err = conn.Ping(); err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	log.Println("Initializing DB connection")

	return conn, nil
}
