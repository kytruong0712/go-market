//go:generate go run github.com/99designs/gqlgen generate

package public

import (
	"github.com/kytruong0712/go-market/product-service/api/internal/controller/categories"

	"github.com/99designs/gqlgen/graphql"
)

func NewSchema(
	categoryCtrl categories.Controller,
) graphql.ExecutableSchema {
	cfg := Config{
		Resolvers: &resolver{
			categoryCtrl: categoryCtrl,
		},
	}

	return NewExecutableSchema(cfg)
}

type resolver struct {
	categoryCtrl categories.Controller
}

// Query returns the QueryResolver
func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct {
	*resolver
}
