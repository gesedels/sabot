package clui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNeed(t *testing.T) {
	// success
	err := Need([]string{"argument"}, 1)
	assert.NoError(t, err)

	// error - not enough arguments
	err = Need(nil, 1)
	assert.EqualError(t, err, `not enough arguments`)
}

func TestSplit(t *testing.T) {
	// success - zero arguments
	name, elems := Split(nil)
	assert.Empty(t, name)
	assert.Empty(t, elems)

	// success - one argument
	name, elems = Split([]string{"name"})
	assert.Equal(t, "name", name)
	assert.Empty(t, elems)

	// success - multiple arguments
	name, elems = Split([]string{"name", "argument"})
	assert.Equal(t, "name", name)
	assert.Equal(t, []string{"argument"}, elems)
}
