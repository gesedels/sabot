package errs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	// setup
	Messages["test"] = "test %s"

	// success
	err := Error("test", "error")
	assert.EqualError(t, err, `test error`)
}
