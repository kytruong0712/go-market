package category

import (
	"context"

	"github.com/kytruong0712/go-market/product-service/api/internal/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// GetCategories gets category items which available as navigation item
	GetCategories(context.Context, GetCategoriesInput) ([]model.Category, error)
}

type impl struct {
	dbConn boil.ContextExecutor
}

// New returns an implementation instance satisfying Repository
func New(dbConn boil.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}
