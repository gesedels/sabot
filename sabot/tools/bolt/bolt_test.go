package bolt

import (
	"testing"

	"github.com/gesedels/sabot/sabot/tools/test"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	// success - true
	ok := contains([]byte("name"), "NAM")
	assert.True(t, ok)

	// success - false
	ok = contains([]byte("name"), "NOPE")
	assert.False(t, ok)
}

func TestDelete(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	err := Delete(db, "alpha")
	assert.NoError(t, err)

	// confirm
	ok, _ := Exists(db, "alpha")
	assert.False(t, ok)
}

func TestExists(t *testing.T) {
	// setup
	db := test.DB(t)

	// success - true
	ok, err := Exists(db, "alpha")
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - false
	ok, err = Exists(db, "nope")
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	pairs, err := Get(db, "alpha")
	assert.Equal(t, test.MockData["alpha"], pairs)
	assert.NoError(t, err)
}

func TestGetValue(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	data, err := GetValue(db, "alpha", "body")
	assert.Equal(t, "Alpha.\n", data)
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	names, err := List(db)
	assert.Equal(t, []string{"alpha", "bravo"}, names)
	assert.NoError(t, err)
}

func TestMatch(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	names, err := Match(db, "ALPH")
	assert.Equal(t, []string{"alpha"}, names)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	names, err := Search(db, "body", "ALPH")
	assert.Equal(t, []string{"alpha"}, names)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	err := Set(db, "name", map[string]string{"attr": "data"})
	assert.NoError(t, err)

	// confirm
	data, _ := GetValue(db, "name", "attr")
	assert.Equal(t, "data", data)
}
