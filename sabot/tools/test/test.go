// Package test implements unit testing data and functions.
package test

import (
	"fmt"
	"testing"

	"github.com/gesedels/sabot/sabot/tools/errs"
	"github.com/stretchr/testify/assert"
)

// AssertError asserts an error is equal to a errs.Error error.
func AssertError(t *testing.T, err error, name string, elems ...any) bool {
	text := fmt.Sprintf(errs.Messages[name], elems...)
	return assert.EqualError(t, err, text)
}
