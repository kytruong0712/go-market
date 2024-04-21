package scalar

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/volatiletech/sqlboiler/v4/types"
)

type JSONSlice []map[string]types.JSON

// UnmarshalGQL unmarshal JSON, a custom scalar type
// Implements graphql.Marshaler
func (v *JSONSlice) UnmarshalGQL(a interface{}) error {
	fmt.Println("fired to UnmarshalGQL")

	jStr, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return err
	}

	return json.Unmarshal(jStr, a)
}

// MarshalGQL marshal JSON, required by gqlgen
// Implements graphql.Marshaler
func (v JSONSlice) MarshalGQL(w io.Writer) {
	out, _ := json.Marshal(v)
	w.Write(out)
}
