package repository

import (
	"context"
	"database/sql"

	"github.com/kytruong0712/go-market/product-service/api/internal/repository/category"
)

// Registry represents the specification of this pkg
type Registry interface {
	// PingPG will check if the DB connection is alive or not
	PingPG(context.Context) error

	// Category returns the Category repo
	Category() category.Repository
}

// New returns an implementation instance which satisfying Registry
func New(pgConn *sql.DB) Registry {
	return impl{
		category: category.New(pgConn),
		pgConn:   pgConn,
	}
}

type impl struct {
	category category.Repository
	pgConn   *sql.DB
}

// Category returns the Category repo
func (i impl) Category() category.Repository {
	return i.category
}
