package main

import (
	"github.com/kytruong0712/go-market/product-service/api/internal/config/httpserver/gql"
	"github.com/kytruong0712/go-market/product-service/api/internal/controller/categories"
	"github.com/kytruong0712/go-market/product-service/api/internal/controller/system"
	"github.com/kytruong0712/go-market/product-service/api/internal/handler/gql/public"

	"github.com/go-chi/chi/v5"
)

type router struct {
	corsOrigins  []string
	systemCtrl   system.Controller
	categoryCtrl categories.Controller
}

func (rtr router) routes(r chi.Router) {
	r.Group(rtr.public)
}

func (rtr router) public(r chi.Router) {
	const prefix = "/gateway/public"

	r.Handle(prefix+"/graphql", gql.Handler(public.NewSchema(
		rtr.categoryCtrl,
	), true))
}
