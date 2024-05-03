package test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Compare allows comparing expected vs actual excluding ignoreFields and shows difference between both
func Compare(t *testing.T, expected interface{}, actual interface{}, model interface{}, ignoreFields ...string) {
	if diff := cmp.Diff(expected, actual, cmpopts.IgnoreFields(model, ignoreFields...)); diff != "" {
		t.Errorf("\n model mismatched. \n expected: %+v \n got: %+v \n diff: %+v", expected, actual, diff)
		t.FailNow()
	}
}
