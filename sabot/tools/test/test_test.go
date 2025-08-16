package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBool(t *testing.T) {
	// setup
	db := MockDB(t)

	// success
	ok := GetBool(db, "select 1")
	assert.True(t, ok)
}

func TestGetInt(t *testing.T) {
	// setup
	db := MockDB(t)

	// success
	data := GetInt(db, "select 1")
	assert.Equal(t, 1, data)
}

func TestGetMap(t *testing.T) {
	// setup
	db := MockDB(t)

	// success
	data := GetMap(db, "select * from Notes where id=1 limit 1")
	assert.Equal(t, "alpha", data["name"])
}

func TestMockDB(t *testing.T) {
	// success
	db := MockDB(t)
	assert.NotNil(t, db)
}
