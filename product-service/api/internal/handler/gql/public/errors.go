package public

import (
	"net/http"

	"github.com/kytruong0712/go-market/product-service/api/internal/config/httpserver"
)

const (
	ErrCodeValidationFailed = "validation_failed"
)

var (
	WebErrHierarchyLevelInvalid = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "invalid hierarchy level"}
)
