package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestDB(t *testing.T) {
	// success
	db := MockDB(t)
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
