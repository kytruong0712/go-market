package system

import (
	"context"
	"errors"
	"net/http"

	"github.com/kytruong0712/go-market/product-service/api/internal/config/httpserver"
)

// CheckDBReady checks for database readiness
func (h Handler) CheckDBReady() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		err := h.systemCtrl.CheckDBReadiness(r.Context())
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}

			return err
		}
		w.Write([]byte("connected!"))

		return nil
	})
}
