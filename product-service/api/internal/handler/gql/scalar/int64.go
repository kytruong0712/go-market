package scalar

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/kytruong0712/go-market/product-service/api/internal/config/httpserver"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalInt64 marshals int64 to string
func MarshalInt64(t int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(strconv.FormatInt(t, 10)))
	})
}

// UnmarshalInt64 unmarshals string to int64
func UnmarshalInt64(v interface{}) (int64, error) {
	var err error
	switch v := v.(type) {
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, &httpserver.Error{
				Status: http.StatusBadRequest,
				Code:   "invalid_gql_field_value",
				Desc:   "Unable to convert value to int64",
			}
		}
		return i, nil
	default:
		err = &httpserver.Error{
			Status: http.StatusBadRequest,
			Code:   "invalid_gql_field_type",
			Desc:   fmt.Sprintf("int64 must be of string type, but got %s instead", reflect.TypeOf(v)),
		}
	}
	return 0, err
}
