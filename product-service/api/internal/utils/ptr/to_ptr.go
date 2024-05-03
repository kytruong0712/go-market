package ptr

import "time"

func ToIntPtr[E int64 | int32 | int16 | int](v E) *E {
	return valToPtr(v)
}

func ToStringPtr(v string) *string {
	return valToPtr(v)
}

func ToTimePtr(v time.Time) *time.Time {
	return valToPtr(v)
}

func ToBoolPtr(v bool) *bool {
	return valToPtr(v)
}

func valToPtr[V any](v V) *V {
	return &v
}
