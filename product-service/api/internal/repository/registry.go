package repository

import (
	"context"
	"database/sql"
)

// Registry represents the specification of this pkg
type Registry interface {
	// PingPG will check if the DB connection is alive or not
	PingPG(context.Context) error
}

// New returns an implementation instance which satisfying Registry
func New(pgConn *sql.DB) Registry {
	return impl{
		pgConn: pgConn,
	}
}

type impl struct {
	pgConn *sql.DB
}
