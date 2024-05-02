package categories

import (
	"context"

	"github.com/kytruong0712/go-market/product-service/api/internal/model"
	"github.com/kytruong0712/go-market/product-service/api/internal/repository"
)

// Controller provides the specification of the functionality provided by this pkg
type Controller interface {
	// GetCategoriesHierarchy gets categories hierarchy
	GetCategoriesHierarchy(context.Context) (model.CategoriesHierarchy, error)
}

type impl struct {
	repo repository.Registry
}

// New returns an implementation instance satisfying Controller
func New(repo repository.Registry) Controller {
	return impl{repo: repo}
}
