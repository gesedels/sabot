package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestDB(t *testing.T) {
	// success
	db := DB(t)
	db.View(func(tx *bbolt.Tx) error {
		for name, pairs := range MockData {
			buck := tx.Bucket([]byte(name))
			assert.NotNil(t, buck)

			for attr, want := range pairs {
				data := buck.Get([]byte(attr))
				assert.Equal(t, want, string(data))
			}
		}

		return nil
	})
}

func TestHas(t *testing.T) {
	// setup
	db := DB(t)

	// success - true
	ok := Has(db, "alpha")
	assert.True(t, ok)

	// success - false
	ok = Has(db, "nope")
	assert.False(t, ok)
}

func TestGet(t *testing.T) {
	// setup
	db := DB(t)

	// success
	data := Get(db, "alpha", "body")
	assert.Equal(t, "Alpha.\n", data)
}

func TestSet(t *testing.T) {
	// setup
	db := DB(t)

	// success
	Set(db, "name", "attr", "data")
	data := Get(db, "name", "attr")
	assert.Equal(t, "data", data)
}
