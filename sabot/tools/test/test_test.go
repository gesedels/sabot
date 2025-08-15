package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// setup
	db := MockDB()

	// success
	data := Get(db, "select * from Notes where id=1 limit 1")
	assert.Equal(t, "alpha", data["name"])
}

func TestMockDB(t *testing.T) {
	// success
	db := MockDB()
	assert.NotNil(t, db)
}
