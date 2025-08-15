package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBool(t *testing.T) {
	// setup
	db := MockDB()

	// success
	ok := GetBool(db, "select 1")
	assert.True(t, ok)
}

func TestGetMap(t *testing.T) {
	// setup
	db := MockDB()

	// success
	data := GetMap(db, "select * from Notes where id=1 limit 1")
	assert.Equal(t, "alpha", data["name"])
}

func TestMockDB(t *testing.T) {
	// success
	db := MockDB()
	assert.NotNil(t, db)
}
