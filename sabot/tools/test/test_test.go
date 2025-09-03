package test

import (
	"testing"

	"github.com/gesedels/sabot/sabot/tools/errs"
	"github.com/stretchr/testify/assert"
)

func TestAssertError(t *testing.T) {
	// setup
	errs.Messages["test"] = "test %s"
	err := errs.Error("test", "error")

	// success
	ok := AssertError(t, err, "test", "error")
	assert.True(t, ok)
}
