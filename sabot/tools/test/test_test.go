package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertSQL(t *testing.T) {
	// setup
	db := MockDB(t)

	// success
	AssertSQL(t, db, "select 123", 123)
}

func TestMockDB(t *testing.T) {
	// success
	db := MockDB(t)
	assert.NotNil(t, db)
}
