package clui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// success - with argument
	amap, err := Parse([]string{"name"}, []string{"argument"})
	assert.Equal(t, map[string]string{"name": "argument"}, amap)
	assert.NoError(t, err)

	// success - with default
	amap, err = Parse([]string{"name:default"}, nil)
	assert.Equal(t, map[string]string{"name": "default"}, amap)
	assert.NoError(t, err)

	// error - argument not provided
	amap, err = Parse([]string{"name"}, nil)
	assert.Nil(t, amap)
	assert.EqualError(t, err, `cannot parse argument "name" - not provided`)
}

func TestSplit(t *testing.T) {
	// success - no arguments
	name, elems := Split(nil)
	assert.Empty(t, name)
	assert.Empty(t, elems)

	// success - one argument
	name, elems = Split([]string{"name"})
	assert.Equal(t, "name", name)
	assert.Empty(t, elems)

	// success - multiple argument
	name, elems = Split([]string{"name", "argument"})
	assert.Equal(t, "name", name)
	assert.Equal(t, []string{"argument"}, elems)
}
