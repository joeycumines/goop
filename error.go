package goop

import (
	"fmt"
)

// OptimizeError is returned from [Model.Optimize], in situations where an
// error code is available.
type OptimizeError struct {
	Code    int
	Message string
}

func (x *OptimizeError) Error() string {
	return fmt.Sprintf(
		"[Code = %d] %s",
		x.Code,
		x.Message,
	)
}

// Is compares err by its code, or if it's a nil pointer. Note nil interface
// value will always return false.
func (x *OptimizeError) Is(err error) bool {
	if x == nil {
		return err == x
	}
	if v, _ := err.(*OptimizeError); v != nil && x.Code == v.Code {
		return true
	}
	return false
}
