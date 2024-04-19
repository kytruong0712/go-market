package main

import (
	"github.com/go-chi/chi/v5"

	"github.com/kytruong0712/go-market/product-service/api/internal/controller/system"
)

type router struct {
	corsOrigins []string
	systemCtrl  system.Controller
}

func (rtr router) routes(r chi.Router) {
	// TODO: add routes here
}
